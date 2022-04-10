import json
import sys
from time import sleep
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
usr = []
nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
          'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
igual = False

if __name__ == "__main__":
    with open('dadosVeiculares.json', 'r', encoding='utf-8') as arq:
        arq_json = json.loads(arq.read())

    for i in range(0, len(arq_json["Veiculo"])):
        #print(f'[ {i} ] {arq_json["Veiculo"][i]["Marca"]} - {arq_json["Veiculo"][i]["Versao"]} - {arq_json["Veiculo"][i]["Modelo"]}\n')
        sleep(0.2)
        sys.stderr.write(
            f'[ {i} ] {arq_json["Veiculo"][i]["Marca"]} - {arq_json["Veiculo"][i]["Versao"]} - {arq_json["Veiculo"][i]["Modelo"]}\n')
        sys.stderr.write('----------------------------------\n')
    valor = input('Qual a marca do seu carro? ')

    while type(valor) != int:
        try:
            valor = int(valor)
        except:
            print('VALOR INVÁLIDO!')
            valor = input('Qual a marca do seu carro? ')

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

    usr.append(valor)
    usr.append(placa)
    print(type(usr), usr)

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    print("Checando instalação de arquivo fbpki:")
    response = loop.run_until_complete(c_hlf.query_installed_chaincodes(
        requestor=admin,
        peers=[callpeer]
    ))
    print(response)

    c_hlf.new_channel(channel_name)

    response = loop.run_until_complete(c_hlf.chaincode_invoke(
        requestor=admin,
        channel_name=channel_name,
        peers=[callpeer],
        cc_name=cc_name,
        cc_version=cc_version,
        fcn='registrarBanco',
        args=[arq_json],
        cc_pattern=None))

    print("Veiculo registrado com sucesso !")