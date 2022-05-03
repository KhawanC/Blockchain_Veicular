import json
from typing import Counter

placas = ['abc1d23', 'dcb3a21', 'abc1234']
contar_linhas = Counter()

with open('teste.json', 'r', encoding='utf-8') as arq_r:
    info = json.loads(arq_r.read())
    
qtd_veiculos_json = len(info["Placas"])

for i in range(len(placas)):
    info["Placas"][qtd_veiculos_json] = placas[i]
    qtd_veiculos_json = len(info["Placas"])

data = json.dumps(info, indent=2)
    
with open('teste.json', 'w', encoding='utf-8') as arq_w:
    arq_w.write(data)
