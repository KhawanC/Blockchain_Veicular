from crypt import methods
from hfc.fabric import Client as client_fabric
import asyncio
from flask import *
from tornado.platform.asyncio import AnyThreadEventLoopPolicy

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"

app = Flask(__name__)

asyncio.set_event_loop_policy(AnyThreadEventLoopPolicy())

@app.route('/inserir_teste/&=<hash>&=<vin>&=<co>')
def hello_word(hash, vin, co):
    return "teste {} {} {}".format(hash, vin, co)


@app.route('/inserir_veiculo/&=<string:hash>&=<string:vin>&=<int:carbono>', methods=['POST'])
def insertBlockchain(hash, vin, carbono):

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    
    callpeer = "peer0." + domain
    
    c_hlf.new_channel(channel_name)

    response = loop.run_until_complete(
        c_hlf.chaincode_invoke(requestor=admin,
                               channel_name=channel_name,
                               peers=[callpeer],                               
                               args=[vin, hash, str(carbono)],
                               cc_name=cc_name,
                               cc_version=cc_version,
                               fcn='registrarVeiculo',
                               cc_pattern=None))
    
    print(response)
    
    return {
        "status": 200,
        "mensagem": "Veiculo registrado com sucesso!"
    }

if __name__ == "__main__":
    app.run(debug=True, port=8001, host="0.0.0.0")






