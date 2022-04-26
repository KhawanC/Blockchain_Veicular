import json

with open('dadosVeicularesBase.json', 'r', encoding='utf-8') as arq:
    banco_json = json.loads(arq.read())

'''contador = 0
for i in banco_json["Veiculo"]:
    sleep(0.2)
    sys.stderr.write(f" [ {contador} ] - {banco_json['Veiculo'][i]['Marca']}"
                     f" - "
                     f"{banco_json['Veiculo'][i]['Versao']} - {banco_json['Veiculo'][i]['Modelo']}")
    sys.stderr.write('\n----------------------------------\n')
    contador += 1

tentativa = 0
valor = int(input('Dentre as opções acima, digite a do seu veículo:  '))
while (valor < 0 or valor > len(banco_json["Veiculo"])):
    tentativa += 1
    if tentativa > 2:
        raise Exception("Muitas tentativas, encerrando o programa")
    valor = int(input("Valor incorreto\nDigite novamente: "))
'''
user = 11
contador = 0
for cdg in banco_json["Veiculo"]:
    if user == contador:
        cdgUsuario = cdg
    contador += 1

print(cdgUsuario)




