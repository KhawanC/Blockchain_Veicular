# import json, codecs
# from random import gauss

# with open("dadosVeicularesBase.json", 'r', encoding='utf8') as f:
#     arq_json = json.load(f)

# print(len(arq_json))
# print(arq_json[0].keys())
# for i in range(len(arq_json)):
#     arq_json[i]["Gasolina_Diesel_Eletrico_-_Cidade_(km/l)"] = gauss(
#         arq_json[i]["Gasolina_Diesel_Eletrico_-_Cidade_(km/l)"], 10)

# for i in range(len(arq_json)):
#     arq_json[i]["Gasolina_Diesel_Eletrico_-_Estrada_(km/l)"] = gauss(
#         arq_json[i]["Gasolina_Diesel_Eletrico_-_Estrada_(km/l)"], 10)

# print('dados de emissao atualziados')

# with open('dadosVeicularesAtualizados.json', 'w', encoding='utf8') as arq:
#     json.dump(arq_json, arq, indent=2, separators=(',', ': '), ensure_ascii=False)

nome = "KauaCassiano"
numero = 0
nome2 = "Ki"
if type(numero) == int:
    print("numero é um numero")
if len(nome2) < 3:
    print("nome 2 tem menos que 2 digitos")
print(nome[0:3] + str(numero) + nome2[0:2])