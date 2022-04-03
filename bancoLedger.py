import sys, json
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de" 
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"


if __name__ == "__main__":

    if len(sys.argv) != 2:
        print("É necessário informar o nome do arquivo .json")
        exit(1)

    #try to retrieve the public key
    try:
        with open(sys.argv[1] + ".json", 'r', encoding='utf-8') as f:
            arq_json = f.read()
    except:
        print("Não foi possivel carregar o arquivo json. Digite novamente!")
        exit(1)

    print("Arquivo ", sys.argv[1], "carregado com sucesso")

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    print("Checking if the chaincode fabpki is properly installed:")
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

    print("Successo em registrar seu banco de dados!")