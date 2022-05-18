from flask import Flask
from flask_restplus import Api

class Server():
    def __init__(self, ):
        self.app = Flask(__name__)
        self.api = Api(self.app,
            version = '1.0',
            titulo = 'API Blockchain',
            descricao = 'API para conectar MongoDb/Chaincodes e o FrontEnd',
            doc = '/docs'
            )

    def run(self, ):
        self.app.run(
            debug = True
        )

server = Server()