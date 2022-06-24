import couchdb, asyncio
from random import gauss
from hfc.fabric import Client as client_fabric

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
couch = couchdb.Server()
db = couch['nmi-channel_fabpki']
listaFabricantes = []

if __name__ == "__main__":

    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:4] == "fab-":
            listaFabricantes.append(i)
    
    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    for i in listaFabricantes:
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
                requestor=admin,
                channel_name=channel_name,
                peers=[callpeer],
                cc_name=cc_name,
                cc_version=cc_version,
                fcn='registrarCredito',
                args=[i],
                cc_pattern=None))
                
    print("Saldo de carbono atualizado com sucesso")