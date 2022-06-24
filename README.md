# 1. Instalando dependências

É recomendado que utilize o **Ubuntu 20.04 LTS**, caso utilize sistema execute o [script de instalação](dependencias/installFabric.sh)

```console
./installFabric.sh
```

**Observação**: Não é necessário rodar o script como *sudo*, quando necessário ele pedirá permissão juntamente de sua senha.

# 2. Gerando MSP artifacts

Primeiramente é necessário checkar a variável FABRIC_CFG_PATH, descomente-a em [configmsp.sh](blockchain/configmsp.sh)

```console
export FABRIC_CFG_PATH=$PWD
```

Após isso, execute o script:

```console
./configmsp.sh
```

Esse script gera os certificados MSP na pasta [crypto-config](blockchain/crypto-config/) que deverão ser devidamente substituídos no [ptb.de.json](blockchain/fabpki-cli/ptb.de.json) e no [inmetro.br.json](blockchain/fabpki-cli/inmetro.br.json) nas linhas 38/39/51

# 3. Gerenciando os containers docker

Vamos então checkar os contéudo na pasta [.env](blockchain/.env), acesse-á e altere de acordo com seu endereço IP. Também devemos acessar [ptb.de.json](blockchain/fabpki-cli/ptb.de.json) e o [inmetro.br.json](blockchain/fabpki-cli/inmetro.br.json) e alterar as linhas 58/76/77 com seu devído endereço IP, mantendo a porta padrões *(:7050 / :7051 / :7053)*.

Para iniciar o container de uma organização em específico utilize o comando: 

```console
docker-compose -f peer-ptb.yaml up -d
```

Se sua organização é responsável por hostear o serviço orderer você também precisará iniciá-la junto de sua organização, informando ambos arquivos .yaml:

```console
docker-compose -f peer-orderer.yaml -f peer-ptb.yaml up -d
```

Caso seja necessário desativar ou restartar os containers utilize os seguintes comandos: 

Desativar:
```console
docker-compose -f peer-ptb.yaml stop
```

Resetar:
```console
./teardown.sh peer-ptb.yaml
```

# 4. Criar a conexão com o Fabric e conectar os peers

Para fazer essa conexão execute o seguinte comando, alterando apenas a organização se necessário:

```console
./configchannel.sh ptb.de -c
```

Em caso de sucesso o Hyperledger Fabric estará rodando em seu servidor, com uma instância da sua organização no sistema. Para ver mais informações do container execute:

```console
docker ps
docker stats
```

# 5. Instalando e instânciando o chaincode

Insira o chaincode dentro da pasta [fabpki](blockchain/fabpki/), volte para a pastar raiz do projeto e execute o seguinte comando para instalar o chaincode:

```console
./configchaincode.sh install cli0 fabpki 1.0
```

Para instânciar o chaincode em toda a rede execute o comando: 

```console
./configchaincode.sh instantiate cli0 fabpki 1.0
```

# 6. Mais algumas dependências


Vamos instalar mais algumas dependências ecessárias para executar os próximos códigos:

```console
cd $HOME
git clone https://github.com/hyperledger/fabric-sdk-py.git
cd fabric-sdk-py
git checkout tags/v0.8.0
sudo make install
```

```console
cd $HOME
git clone https://github.com/hyperledger/fabric-sdk-py.git
cd fabric-sdk-py
git checkout tags/v0.8.0
sudo make install
```

# 7. Iniciando API

Execute a API [requisicaoLedger.py](API/requisicaoLedger.py) com:

```console
python3 requisicaoAPI.py
```

Em seu terminal aparecerá dois endereços IP. Geralmente em ambiente local é utilizado o *127.0.0.1 (o mesmo que localhost)*, entretando, caso queira fazer requisições fora do seu ambiente local deverá utilizar o IP da sua rede.

# 8. Fazendo requisições 

Instale o Insomnia ou qualquer outro ambiente para testes de requisições HTTPS, caso utilize o Insomnia importe o arquivo [backup-insomnia.json](extras/backup_insomnia.json), caso você não possua o Insomia é possível verificar o corpo das requisições apenas pelo arquivo JSON.


### POST Veiculo
![Post Veiculo](https://i.imgur.com/pTDURzy.png)

Esse é o corpo necessário para inserir um veiculo na rede, perceba que em existe um parâmetro chamado "Vin", ele é super importante pois é necessário para o registro dos fabricantes dentro da rede, então se certifique de que esse Vin seja um parâmetro válido para evitar conflitos. 

### POST Fabricante
![Post Fabricante](https://i.imgur.com/5b8Q4QL.png)

A princípio não será necessário utilizar esse endpoint, pois os fabricantes são altomaticamente inseridos ao verificarem o fabricante dos veiculos, isso é feito para que os créditos do veículo não se percam caso o fabricante do veículo não exista na rede.

### POST Saldo
![Post Saldo](https://i.imgur.com/1dAMC2T.png)

Esse endpoint é utilizado para que a quantidade de Co2 acumulado dos fabricantes seja convertido em saldo com base em uma meta, em questão de sistema esse endpoint deve ser ativado em períodos equidistantes para uma amostragem mais precisa desses cálculos.

# 9. Leilão

Após o cálculo do saldo dos fabricantes será possível o início das transações dentro do sistema, onde fabricantes com saldo de carbono negativo podem comprar o saldo de outros fabricantes com uma moeda fiduciário. Todos começam com 10.000 dessa moeda fiduciária para testes.

## POST Iniciar Ordem
![Post Iniciar Ordem](https://i.imgur.com/1dAMC2T.png)

Com esse endpoint é possível iniciar uma nova *ordem*. Como parâmetro temos o proprietário da ordem, que virá com o id do fabricante dentro dos registros do blockchain *(Nessa rede todo o fabricante inicia com "fab-")*. Temos o tipo da transação, sendo permitido preencher com "comprar" e "vender". Por último o saldo ofertado, caso seja preenchido "comprar" no campo de tipo de transação esse saldo representará o saldo fiduciário, caso preenchido "vender" representará o saldo de carbono

## POST Registrar Lance
![Post Lance](https://i.imgur.com/4adFlzR.png)

Esse endpoint representa um lance dentro da ordem criada anteriormente. Ela tem no corpo do seu JSON o id da transação criada *(Existe um endpoint de GET para as transações no JSON de backup do Insomnia)*. O valor do seu lance, existindo os detalhes para caso a ordem seja de venda venda ou compra explicados no item anterior. Por último o id do comprador.

Existem algumas regras de negócio dentro do sistema, como a impossibilidade de fazer dois lances seguidos para o mesmo fabricante, ou até fazer um lance não possuindo saldo suficiente. Se atente com isso!

## POST Fechar ordem
![Post Fechar Ordem](https://i.imgur.com/4adFlzR.png)

Nesse último endpoint será possível fechar a ordem para novos lances, em seu fechamento o respectivo saldo e computado para o proprietario e o ultimo lancista da ordem, essa ordem também ficará fechada para novos lances.