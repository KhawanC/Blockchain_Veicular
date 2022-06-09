import json, sys, random, couchdb
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
couch = couchdb.Server()
server = couchdb.Server('http://localhost:5984/_utils')
db = couch['nmi-channel_fabpki']
listaModelos = []
listaPlacas = []

# Variáveis para verificação da placa
nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
          'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

if __name__ == "__main__":

    #Loop para criar 10 placas
    for i in range(0, 50):
        
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
        listaPlacas.append(placaCompleto)
    
    #Iniciando função para verificar se há duplicadas
    removerDuplicados(listaPlacas)
    
    #Acesso couchdb e recuperando modelos
    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:6] == "model-":
            listaModelos.append(i)
    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    #Loop para enviar ao chaincode cada placa 
    for i in range(len(listaPlacas)):
        x = random.randint(0, (len(listaModelos)-1))
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarVeiculo',
            args=[listaPlacas[i], listaModelos[x]],
            cc_pattern=None))
        
    print("Veiculos registrado com sucesso !")
