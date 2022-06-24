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
    
    for i in range(0, len(listaTransacoes)):
        print(str(i) + ' - ' + listaTransacoes[i])
    
    idTransacao = int(input('Esses são as transações, qual você deseja? '))       
    
    compradorInfor = {}
    
    for doc in db.find({
            "selector": {
                "_id": "{id}".format(id=listaTransacoes[idTransacao])
            }}):
        query_info = json.dumps(doc, indent=4, sort_keys=True)
        query_json = json.loads(query_info)
        transacaoInfor = query_json
    
    print('\nInformações: ')
    print('Id Comprador: ' + transacaoInfor['IdComprador'])
    print('Propietário da Ordem: ' + transacaoInfor['ProprietarioOrdem'])
    print('Status da Ordem: ' + transacaoInfor['StatusOrdem'])
    print('Tipo da Transação: ' + transacaoInfor['TipoTransacao'])
    print('Saldo Ofertado: ' + transacaoInfor['SaldoOfertado'])
    print('Ultimo lance: ' + transacaoInfor['ValorLance'])
    
    lance = input('\nQual será o seu lance ? ')
    
    for i in range(0, len(listaFabricantes)):
        print(str(i) + ' - ' + listaFabricantes[i])
        
    idComprador = int(input('\nQuem é você? '))
    
    print('Dados finais: ')
    print('\n' + listaFabricantes[idComprador] + '\n' + listaTransacoes[idTransacao] + '\n' + lance)
    
    
    response = loop.run_until_complete(c_hlf.chaincode_invoke(
                requestor=admin,
                channel_name=channel_name,
                peers=[callpeer],
                cc_name=cc_name,
                cc_version=cc_version,
                fcn='ordemLance',
                args=[listaTransacoes[idTransacao], lance, listaFabricantes[idComprador]],
                cc_pattern=None))
        