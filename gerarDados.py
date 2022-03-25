from random import gauss
import numpy as np, json
import math, openpyxl as exl

def ler_json(arq_json):
    with open(arq_json, 'r', encoding='utf-8') as f:
        return json.load(f)

tabela = exl.load_workbook(filename= 'lista_carros.xlsx')
veiculos = ler_json('dados_veiculares.json')
tabela_act = tabela.active
sigma = 10

user = []
emissoes = [107, 112, 101, 98, 97]

placa = input('Insira sua placa: ')
opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
while int(opt) == 0:
    print('-='*5)
    placa = input('Insira sua placa: ')
    opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
    while int(opt) < 0 or int(opt) > 1:
        print('-='*5)
        print('O valor deve ser entre 1 e 2')
        opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')

#for i in range(0, 3):


for cel in tabela_act['F']:
    try:
        aleat_cel = gauss(float(cel.value),sigma)
        linha = cel.row
        tabela_act[f'F{linha}'] = aleat_cel
    except:
        print('error')

tabela.save('lista_carros(2).xlsx')