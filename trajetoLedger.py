import requests
from hfc.fabric import Client as client_fabric
import asyncio
from time import sleep

bingMapsKey = 'AsmuM7jTKB5hKGiXpA15vY2toMIYZTEq4G_KgLBie0M-BzAeOE17ntxmg4fmC30Q'
endUsr = []
endFinal = []
domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
nums = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
letras = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
          'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

if __name__ == "__main__":
    placa = input("Insira sua placa: ").upper()

    # Leitura da placa e verificação de conformidade da mesma
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

    # Loop que pergunta o endereço de partida do usuário
    for i in range(0, 4):
        if i == 0:
            info = input('Digite a sua rua: ')
        if i == 1:
            info = input('Digite o seu bairro: ')
        if i == 2:
            info = input('Digite a sua cidade: ')
        if i == 3:
            info = input('Digite o seu estado: ')
    endUsr.append(info)

# Simulador de loading...
    for i in range(0, 2):
        print('.')
        sleep(0.5)
    print('Agora vamos verificar o seu destino...')
    sleep(1)
    for i in range(0, 2):
        print('.')
        sleep(0.5)

# Loop que pergunta endereço de destino do usuário
    for i in range(0, 4):
        if i == 0:
            info = input('Digite a sua rua: ')
        if i == 1:
            info = input('Digite o seu bairro: ')
        if i == 2:
            info = input('Digite a sua cidade: ')
        if i == 3:
            info = input('Digite o seu estado: ')
    endFinal.append(info)

# Endereço do trajeto inicial e final juntos em uma única string
    endJuntoUsr = ' '.join(str(e) for e in endUsr)
    endJuntoFinal = ' '.join(str(e) for e in endFinal)

# Request para consultar endereço no sistema Bing Maps
    rotaUrl = "http://dev.virtualearth.net/REST/V1/Routes/Driving?wp.0=" + \
        endJuntoUsr + '&wp.1=' + endJuntoFinal+"/&key=" + bingMapsKey
    jsonRequest = requests.get(url=rotaUrl)
    resultado = jsonRequest.json()
    distancia = resultado["resourceSets"][0]["resources"][0]["travelDistance"]

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    response = loop.run_until_complete(c_hlf.chaincode_invoke(
        requestor=admin,
        channel_name=channel_name,
        peers=[callpeer],
        cc_name=cc_name,
        cc_version=cc_version,
        fcn='registrarUsuario',
        args=[],
        cc_pattern=None))
