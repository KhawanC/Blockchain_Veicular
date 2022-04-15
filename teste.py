import json
from random import gauss

argumentos = []
data = []

with open("dadosVeicularesBase.json") as f:
    arq_json = json.load(f)

for i in arq_json["Veiculo"]:
    arq_json["Veiculo"][i]["Emissao"] = gauss(
    arq_json["Veiculo"][i]["Emissao"], 10)

for i in arq_json["Veiculo"]:
    data.append(arq_json["Veiculo"][i])
    print(i, arq_json["Veiculo"][i])
    print("-------------------")






