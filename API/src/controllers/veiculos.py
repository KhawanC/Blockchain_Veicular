from flask import Flask
from flask_cors import CORS, cross_origin
from flask_restx import Api, Resource
from src.server.instance import server
import src.controllers.querys


app, api= server.app, server.api
cors = CORS(app)

app.config['CORS_HEADERS'] = 'Content-Type'

@api.route('/')
class rotaPadrao(Resource):
    def get(self, ):
        return "Hello World"

@api.route('/listaVeiculos')
class listaVeiculos(Resource):
   def get(self, ):
      info = src.controllers.querys.queryVeiculos()
      return info
     
@api.route('/listaTrajetos')
class listaTrajetos(Resource):
   def get(self, ):
      info = src.controllers.querys.queryViagens()
      return info