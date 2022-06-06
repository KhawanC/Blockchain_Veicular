import random, couchdb
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
couch = couchdb.Server()
server = couchdb.Server('http://localhost:5984/_utils')
db = couch['nmi-channel_fabpki']
items = []

if __name__ == "__main__":
    
    #Acesso couchdb e recuperando modelos
    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:5] == "user-":
            items.append(i)
            
    loop = asyncio.get_event_loop()

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)
    
    #Fazer um loop para cada veiculo, associando um trajeto entre 0 e 80 para eles
    for i in range(len(items)):
        placa = items[i]
        distancia = random.randint(0,120)
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarTrajeto',
            args=[placa, str(distancia)],
            cc_pattern=None))
    
    print("Trajeto de veiculos registrado com sucesso")
