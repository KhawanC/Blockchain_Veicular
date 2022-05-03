import json, sys, random
from time import sleep
from hfc.fabric import Client as client_fabric
import asyncio
from iteration_utilities import duplicates
from iteration_utilities import unique_everseen

#Função para verificar se há elementos duplicados em uma lista
def removerDuplicados(lista):
    return list(unique_everseen(duplicates(lista)))

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"

# Variáveis para verificação da placa
nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
          'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

#Lista para armazenar todas as placas criadas
bancoPlacas = []

if __name__ == "__main__":

    # Método para ler arquivo json atualizado pelo bancoLedger.py
    with open('dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq:
        banco_json = json.loads(arq.read())

    #Loop para criar 20 placas
    for i in range(0, 20):
        
        #Criar uma lista para inserir as letras/numeros no padrão Mercosul (AAA1A11)
        placa = []

        #Criando as primeiras 3 letras aleatóriamente
        for i in range(0, 3):
            randLetra1 = random.randint(0, len(letras))
            if(randLetra1 == 26):
                randLetra1 -= 1
            #Inserindo valores na lista 'placa'
            placa.append(letras[randLetra1])

        #Criando um número aleatório
        randNum1 = random.randint(0, len(nums))
        if(randNum1 == 10):
            randNum1 -= 1
        placa.append(str(nums[randNum1]))

        #Criando mais uma letra de forma aleatória
        randLetra2 = random.randint(0, len(letras))
        if(randLetra2 == 26):
            randLetra2 -= 1
        placa.append(letras[randLetra2])

        #Criando os ultimos 2 números
        for i in range(0, 2):
            randNum2 = random.randint(0, len(nums))
            if(randNum2 == 10):
                randNum2 -= 1
            placa.append(str(nums[randNum2]))

        #Concatenando todos os valores da lista em uma única String
        placaCompleto = "".join(placa)
        
        #Inserindo a placa na lista de placas 
        bancoPlacas.append(placaCompleto)
    
    #Iniciando função para verificar se há duplicadas
    removerDuplicados(bancoPlacas)

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    for i in range(len(bancoPlacas)):
        valor = random.randint(0, len(banco_json["Modelo_Veiculos"]) - 1)
        contador2 = 0
        for cdg in banco_json["Modelo_Veiculos"]:
            if valor == contador2:
                cdgUsuario = cdg
            contador2 += 1
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarUsuario',
            args=[bancoPlacas[i], cdgUsuario],
            cc_pattern=None))
        
    qtd_veiculos_json = len(banco_json["Placas"])

    for i in range(len(bancoPlacas)):
        banco_json["Placas"][qtd_veiculos_json] = bancoPlacas[i]
        qtd_veiculos_json = len(banco_json["Placas"])

    data = json.dumps(banco_json, indent=2)
    
    with open('dadosVeicularesAtualizados.json', 'w', encoding='utf-8') as arq_w:
        arq_w.write(data)
        
    print("Veiculo registrado com sucesso !")
