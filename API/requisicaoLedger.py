from crypt import methods
from hfc.fabric import Client as client_fabric
from flask import *
from tornado.platform.asyncio import AnyThreadEventLoopPolicy
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"

app = Flask(__name__)

asyncio.set_event_loop_policy(AnyThreadEventLoopPolicy())

@app.route('/inserir_veiculo', methods=['POST'])
def insertVeiculo(hash, vin, carbono):
    
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
                               args=[request_data["Vim"], request_data["Hash"], request_data["Co2"]],
                               cc_name=cc_name,
                               cc_version=cc_version,
                               fcn='registrarVeiculo',
                               cc_pattern=None))
    
    print(response)
    
    return {
        "status": 200,
        "mensagem": "Veiculo registrado com sucesso!"
    }

@app.route('/inserir_fabricante', methods=['POST'])
def insertFabricante():
    
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
                               args=[request_data["nome"]],
                               cc_name=cc_name,
                               cc_version=cc_version,
                               fcn='registrarFabricante',
                               cc_pattern=None))
    
    print(response)
    
    return {
        "status": 200,
        "mensagem": "Fabricante registrado com sucesso!"
    }
    

if __name__ == "__main__":
    app.run(debug=True, port=8001, host="0.0.0.0")






