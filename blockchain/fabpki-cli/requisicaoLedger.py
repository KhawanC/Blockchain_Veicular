from crypt import methods
from hfc.fabric import Client as client_fabric
from flask import *
from tornado.platform.asyncio import AnyThreadEventLoopPolicy
from vininfo import Vin
import asyncio, couchdb, json

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
import couchdb, json

app = Flask(__name__)

asyncio.set_event_loop_policy(AnyThreadEventLoopPolicy())    

@app.route('/fabricante', methods=['POST', 'GET'])
def Fabricante():
    if request.method == 'GET':
        listaFabricantes = []
        server = couchdb.Server('http://localhost:5984/_utils')
        couch = couchdb.Server()
        db = couch['nmi-channel_fabpki']
        
        for doc in db.view('_all_docs'):
                i = doc['id']
                if i[0:4] == "fab-":
                    for doc in db.find({
                            "selector": {
                            "_id": "{id}".format(id=i)
                            }}):
                            query_info = json.dumps(doc, indent=4, sort_keys=True)
                            query_json = json.loads(query_info)
                            infoVeiculo = query_json
                            listaFabricantes.append(infoVeiculo)
                    
        return json.dumps(listaFabricantes), 200    
    
    if request.method == 'POST':
        
        request_data = request.get_json()

        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        
        callpeer = "peer0." + domain
        
        c_hlf.new_channel(channel_name)
        
        fab_nome = request_data["nome"]

        response = loop.run_until_complete(
            c_hlf.chaincode_invoke(requestor=admin,
                                channel_name=channel_name,
                                peers=[callpeer],                               
                                args=[fab_nome.upper()],
                                cc_name=cc_name,
                                cc_version=cc_version,
                                fcn='registrarFabricante',
                                cc_pattern=None))
        
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Fabricante registrado com sucesso"}), status=201, mimetype='application/json')
    
@app.route('/veiculo', methods=['GET', 'POST'])
def Veiculo():
    if request.method == 'GET':
        listaVeiculos = []
        server = couchdb.Server('http://localhost:5984/_utils')
        couch = couchdb.Server()
        db = couch['nmi-channel_fabpki']
        
        for doc in db.view('_all_docs'):
                i = doc['id']
                if i[0:5] == "veic-":
                    for doc in db.find({
                            "selector": {
                            "_id": "{id}".format(id=i)
                            }}):
                            query_info = json.dumps(doc, indent=4, sort_keys=True)
                            query_json = json.loads(query_info)
                            infoVeiculo = query_json
                            listaVeiculos.append(infoVeiculo)
                    
        return json.dumps(listaVeiculos), 200
    
    if request.method == 'POST':
        
        request_data = request.get_json()
        
        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        
        callpeer = "peer0." + domain
        
        c_hlf.new_channel(channel_name)
        
        vin = Vin(request_data["Vim"])
        fabNome = (vin.manufacturer).upper().replace(" ", "")

        response = loop.run_until_complete(
            c_hlf.chaincode_invoke(requestor=admin,
                                channel_name=channel_name,
                                peers=[callpeer],                               
                                args=[request_data["Vim"], request_data["Hash"], request_data["Co2"], fabNome],
                                cc_name=cc_name,
                                cc_version=cc_version,
                                fcn='registrarVeiculo',
                                cc_pattern=None))
        
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Veiculo registrado com sucesso"}), status=201, mimetype='application/json')

@app.route('/saldo', methods=['POST'])
def Saldo():
    if request.method == 'POST':
        listaVeiculos = []
        server = couchdb.Server('http://localhost:5984/_utils')
        couch = couchdb.Server()
        db = couch['nmi-channel_fabpki']
        loop = asyncio.get_event_loop()
        c_hlf = client_fabric(net_profile=(domain + ".json"))
        admin = c_hlf.get_user(domain, 'Admin')
        callpeer = "peer0." + domain
        c_hlf.new_channel(channel_name)
        
        for doc in db.view('_all_docs'):
                i = doc['id']
                if i[0:4] == "fab-":
                    listaVeiculos.append(i)
        
        for i in listaVeiculos:
            response = loop.run_until_complete(
                c_hlf.chaincode_invoke(requestor=admin,
                                    channel_name=channel_name,
                                    peers=[callpeer],                               
                                    args=[i],
                                    cc_name=cc_name,
                                    cc_version=cc_version,
                                    fcn='registrarCredito',
                                    cc_pattern=None))
            
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Saldo dos fabricantes atualizados com suceso"}), status=201, mimetype='application/json')
            

@app.route('/ordem', methods=['GET', 'POST'])
def Ordem():
    if request.method == 'GET':
        listaTransacoes = []
        server = couchdb.Server('http://localhost:5984/_utils')
        couch = couchdb.Server()
        db = couch['nmi-channel_fabpki']
        
        for doc in db.view('_all_docs'):
                i = doc['id']
                if i[0:6] == "trans-":
                    for doc in db.find({
                            "selector": {
                            "_id": "{id}".format(id=i)
                            }}):
                            query_info = json.dumps(doc, indent=4, sort_keys=True)
                            query_json = json.loads(query_info)
                            infoVeiculo = query_json
                            listaTransacoes.append(infoVeiculo)
                    
        return json.dumps(listaTransacoes), 200

    if request.method == 'POST' :
        
        request_data = request.get_json()
        
        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        
        callpeer = "peer0." + domain
        
        c_hlf.new_channel(channel_name)

        response = loop.run_until_complete(
            c_hlf.chaincode_invoke(requestor=admin,
                                channel_name=channel_name,
                                peers=[callpeer],                               
                                args=[request_data["ProprietarioOrdem"], request_data["TipoTransacao"], request_data["SaldoOfertado"]],
                                cc_name=cc_name,
                                cc_version=cc_version,
                                fcn='anunciarOrdem',
                                cc_pattern=None))
        
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Ordem de transação registrada com sucesso"}), status=201, mimetype='application/json')    

@app.route('/lance', methods=['POST'])
def Lance():
    if request.method == 'POST' :
        
        request_data = request.get_json()
        
        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        
        callpeer = "peer0." + domain
        
        c_hlf.new_channel(channel_name)

        response = loop.run_until_complete(
            c_hlf.chaincode_invoke(requestor=admin,
                                channel_name=channel_name,
                                peers=[callpeer],                               
                                args=[request_data["IdTransacao"], request_data["ValorLance"], request_data["IdComprador"]],
                                cc_name=cc_name,
                                cc_version=cc_version,
                                fcn='ordemLance',
                                cc_pattern=None))
        
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Lance registradao com sucesso"}), status=201, mimetype='application/json')    

@app.route('/fechar_ordem', methods=['POST'])
def FecharOrdem():
    if request.method == 'POST' :
        
        request_data = request.get_json()
        
        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        
        callpeer = "peer0." + domain
        
        c_hlf.new_channel(channel_name)

        response = loop.run_until_complete(
            c_hlf.chaincode_invoke(requestor=admin,
                                channel_name=channel_name,
                                peers=[callpeer],                               
                                args=[request_data["IdTransacao"], request_data["IdProprietario"]],
                                cc_name=cc_name,
                                cc_version=cc_version,
                                fcn='fecharOrdem',
                                cc_pattern=None))
        
        return Response(response=json.dumps({
            "status": 201,
            "mensagem": "Lance registradao com sucesso"}), status=201, mimetype='application/json') 

if __name__ == "__main__":
    app.run(debug=True, port=8001, host="0.0.0.0")






