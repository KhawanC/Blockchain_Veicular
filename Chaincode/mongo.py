import couchdb, json

items = []
couch = couchdb.Server()

if __name__ == "__main__":
   with open('dadosVeicularesBase.json', 'r', encoding='utf-8') as arq:
         banco_json = json.loads(arq.read())

   # connect to MongoDB
   server = couchdb.Server('http://localhost:5984/_utils')
   # Acess an existing database
   db = couch['nmi-channel_fabpki']

   for doc in db.view('_all_docs'):
         i = doc['id']
         if i[0:7] == "travel-":
            print(i)
            items.append(i)
            print(len(items))
   print(items[0])


