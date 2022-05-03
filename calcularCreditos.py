import json

mediaEmissao = 0

with open('../fabpki-cli/dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq:
    banco_json = json.loads(arq.read())

for i in banco_json["Veiculo"]:
    if banco_json["Veiculo"][i]["Categoria"] == "Compacto":
        mediaEmissao += (banco_json["Veiculo"][i]["Emissao"]*1.3)
    if banco_json["Veiculo"][i]["Categoria"] == "Medio":
        mediaEmissao += (banco_json["Veiculo"][i]["Emissao"]*1.6)

print(mediaEmissao * 30)