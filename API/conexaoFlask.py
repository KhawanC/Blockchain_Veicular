from flask import Flask, jsonify
from flask_cors import CORS
import mongo

app = Flask(__name__)
CORS(app)

#Data 0 --> ONde ficam os veículos
data = []
data.append(mongo.veiculos)

@app.route('/data', methods=['GET'])
def get():
    return jsonify({"Data": data})

def index():
    return "Entrou na API"

if __name__ == "__main__":
    app.run(host='localhost', port=8989)