from flask import Flask
from flask_restplus import Api, Resource
from src.server.instance import server
import couchdb, json

app, api = server.app, server.api

#Lista para armazenar o JSON dos veículos
veiculos = []

#Variavel para iniciar conexao com o CouchDb
couch = couchdb.Server()

#With Open para abrir o banco de dados local de veículos e armazenar em uma variável
with open('../../fabpki-cli/dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq:
      banco_json = json.loads(arq.read())

# connect to MongoDB
server = couchdb.Server('http://192.168.0.105:5984/_utils')
# Acess an existing database
db = couch['nmi-channel_fabpki']

#Loop para resgatar a informação dos veiculos no COuchDb e appendar na lista de veiculos
for i in banco_json["Modelo_Veiculos"]:
   for doc in db.find(
   {
      "selector": {
         "IdCdgCategoria": "{}".format(i)
      }
   }):
      query_info = json.dumps(doc, indent=4, sort_keys=True)
      query_json = json.loads(query_info)
      veiculos.append(query_json)

@api.route('/listaVeiculos')
class listaVeiculos(Resource):
    def get(self, ):
        return veiculos