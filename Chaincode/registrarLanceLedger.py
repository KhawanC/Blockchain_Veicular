import couchdb, asyncio, json
from hfc.fabric import Client as client_fabric

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
couch = couchdb.Server()
db = couch['nmi-channel_fabpki']
listaTransacoes = []
listaFabricantes = []

if __name__ == "__main__":
    
    loop = asyncio.get_event_loop()
    c_hlf = client_fabric(net_profile=(domain + ".json"))
    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain
    c_hlf.new_channel(channel_name)
    
    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:4] == "fab-":
            listaFabricantes.append(i)
    
    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:6] == "trans-":
            listaTransacoes.append(i)
            
    for doc in db.find({
            "selector": {
                "_id": "trans-KIA9f90dfc0713b8b7e3b22044487c9b091244720f9"
            }}):
        query_info = json.dumps(doc, indent=4, sort_keys=True)
        query_json = json.loads(query_info)
        print(query_json)