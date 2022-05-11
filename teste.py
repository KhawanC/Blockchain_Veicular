import json
import random

placas = ['abc1d23', 'dcb3a21', 'abc1234']

with open('dadosVeicularesAtualizados.json', 'r', encoding='utf-8') as arq_r:
    info = json.loads(arq_r.read())
    
qtd_veiculos_json = len(info["Placas"])

for i in range(qtd_veiculos_json):
    placa = info["Placas"][str(i)]
    print(placa)
        
