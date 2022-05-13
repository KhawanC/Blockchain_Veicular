import couchdb, json

def somar_elementos(lista):
   soma = 0
   for numero in lista:
      soma += numero
   return soma

distanciaLista = []
couch = couchdb.Server()

with open('dadosVeicularesBase.json', 'r', encoding='utf-8') as arq:
      banco_json = json.loads(arq.read())

# connect to MongoDB
server = couchdb.Server('http://192.168.0.105:5984/_utils')
# Acess an existing database
db = couch['nmi-channel_fabpki']

for i in banco_json["Modelo_Veiculos"]:
   for doc in db.find(
   {
      "selector": {
         "IdCdgCategoria": "{}".format(i)
      }
   }):
      query_info = json.dumps(doc, indent=4, sort_keys=True)
      query_json = json.loads(query_info)
      distanciaLista.append(query_json["AcumuladorDistancia"])
   
listInt = list(map(float, distanciaLista))
distAcumulado = sum(listInt)
print(distAcumulado)

