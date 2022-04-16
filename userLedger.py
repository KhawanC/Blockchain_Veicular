import json
import sys
from time import sleep
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"

#Variáveis para verificação da placa
nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
          'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
tentativa = 0
contador = 0
igual = False

if __name__ == "__main__":

    #Método para ler arquivo json atualizado pelo bancoLedger.py
    with open('dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq:
        banco_json = json.loads(arq.read())

    #Loop para acessa os dados json e imprimete todos os veículos ordenadamente
    contador = 0
    for i in banco_json["Veiculo"]:
        sleep(0.2)
        sys.stderr.write(f" [ {contador} ] - {banco_json['Veiculo'][i]['Marca']}"
                         f" - "
                         f"{banco_json['Veiculo'][i]['Versao']} - {banco_json['Veiculo'][i]['Modelo']}")
        sys.stderr.write('\n----------------------------------\n')
        contador += 1

    valor = int(input('Dentre as opções acima, digite a do seu veículo:  '))
    while (valor < 0 or valor > len(banco_json["Veiculo"])):
        tentativa += 1
        if tentativa > 2:
            raise Exception("Muitas tentativas, encerrando o programa")
        valor = int(input("Valor incorreto\nDigite novamente: "))

    tentativa = 0
    while type(valor) != int:
        try:
            valor = int(valor)
        except:
            print('VALOR INVÁLIDO!')
            valor = int(input('Qual a marca do seu carro? '))

    placa = input("Insira sua placa: ").upper()

    if len(placa) != 7:
        raise Exception("PLACA INVÁLIDA")

    try:
        str(placa[0])
        str(placa[1])
        str(placa[2])
        int(placa[3])
        str(placa[4])
        int(placa[5])
        int(placa[6])
    except:
        raise Exception("Placa Inválida")

    for i in range(0, len(letras)):
        if placa[0] == letras[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(letras)):
        if placa[1] == letras[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(letras)):
        if placa[2] == letras[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(nums)):
        if int(placa[3]) == nums[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(letras)):
        if placa[4] == letras[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(nums)):
        if int(placa[5]) == nums[i]:
            igual = True
            break
        else:
            igual = False

    for i in range(0, len(nums)):
        if int(placa[6]) == nums[i]:
            igual = True
            break
        else:
            igual = False

    if igual == False:
        raise Exception("PLACA INVÁLIDA")

    for cdg in banco_json["Veiculo"]:
        if valor == contador:
            cdgUsuario = cdg
        contador += 1

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    print("Checando instalação de arquivo fbpki:")

    c_hlf.new_channel(channel_name)

    response = loop.run_until_complete(c_hlf.chaincode_invoke(
        requestor=admin,
        channel_name=channel_name,
        peers=[callpeer],
        cc_name=cc_name,
        cc_version=cc_version,
        fcn='registrarUsuario',
        args=[placa, cdgUsuario],
        cc_pattern=None))

    print("Veiculo registrado com sucesso !")
