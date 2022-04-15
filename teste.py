import json
from random import gauss

data = []

with open("dadosVeicularesBase.json") as f:
    arq_json = json.load(f)

for i in arq_json["Veiculo"]:
    arq_json["Veiculo"][i]["Emissao"] = gauss(
    arq_json["Veiculo"][i]["Emissao"], 10)


#ideia 1: colocar valores dentro de uma lista para enviar ao ledger
for i in arq_json["Veiculo"]:
    data.append(arq_json["Veiculo"][i])
    print(i, arq_json["Veiculo"][i])
    print("-------------------")

for i in arq_json["Veiculo"]:
    categoria = arq_json["Veiculo"][i]["Categoria"]
    marca =arq_json["Veiculo"][i]["Marca"]
    versao = arq_json["Veiculo"][i]["Versao"]
    modelo = arq_json["Veiculo"][i]["Modelo"]
    emissao = arq_json["Veiculo"][i]["Emissao"]
    print("Categoria: {} // tipo: {}\nMarca: {} tipo \nVersao: \nModelo: {}\nEmissao: {}".format(categoria, type(categoria),marca, versao, modelo, emissao))
    print("--------"*2)







