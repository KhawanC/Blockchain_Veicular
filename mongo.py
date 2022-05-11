import couchdb
import json

couch = couchdb.Server()

# connect to MongoDB
server = couchdb.Server('http://10.25.103.88:5984/_utils')
# Acess an existing database
db = couch['nmi-channel_fabpki']

for id in db:
	print(id)
	
for doc in db.find({
   "selector": {
      "_id": "PPU6G22"
   }
}):
    print(json.dumps(doc, indent=4, sort_keys=True))
   

