import couchdb, json

#FUnção para saber a quantidade de elementos de uma lista
def somar_elementos(lista):
   soma = 0
   for numero in lista:
      soma += numero
   return soma

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
      #distanciaLista.append(query_json["AcumuladorDistancia"])  

      
   #listInt = list(map(float, distanciaLista))
   #distAcumulado = sum(listInt)
   #print(distAcumulado)

