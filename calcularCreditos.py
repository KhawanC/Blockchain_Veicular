import json
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"

if __name__ == "__main__":

    #Ler arquivo json e coloca-lo em uma variavel chamada "info"
    with open('dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq_r:
        info = json.loads(arq_r.read())
    
    #Variavel que armazenará a quantidade de placas dentro do arquivo json
    qtd_veiculos_json = len(info["Placas"])

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)
    
    #Fazer um loop para enviar as placas do json ao smart contract
    for i in range(qtd_veiculos_json):
        placa = info["Placas"][str(i)]
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='calcularEmissao',
            args=[placa],
            cc_pattern=None))
    
    for i in range(qtd_veiculos_json):
        placa = info["Placas"][str(i)]
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='calcularCreditos',
            args=[placa],
            cc_pattern=None))
    
    print("Calculos de créditos de carbono efetuados com sucesso")
