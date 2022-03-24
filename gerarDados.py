from random import gauss
import numpy as np

import math, openpyxl as exl

tabela = exl.load_workbook(filename= 'lista_carros.xlsx')
tabela_act = tabela.active
sigma = 10
opt = 0
qst = ['categoria', 'marca', 'versão', 'modelo']
user = []
emissoes = [107, 112, 101, 98, 97]

while opt == 0:
    placa = input('Insira sua placa: ')
    opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim')
    while opt != 0 or opt != 1:
        print('Valor deve estar entre 1 e 0')
        opt = input('Tem certeza que essa é sua placa?\n [ 0 ] Não\n[ 1 ] Sim')

for i in range(0, 4):
    print('Insira o/a {} do seu veículo'.format(qst[i]))
    if i == 0:
        cat()
        acml = input()
        user.append(acml)
        
    if i == 1 and acml == :
       
    if i == 2 and user[1] == "Volks Wagen":
        
        


for cel in tabela_act['F']:
    try:
        aleat_cel = gauss(float(cel.value),sigma)
        linha = cel.row
        tabela_act[f'F{linha}'] = aleat_cel
    except:
        print('error')

tabela.save('lista_carros(2).xlsx')

def cat():
    print('''[ 0 ] Compacto
    [ 1 ] Médio
    [ 2 ] Grande''')

def marca_comp():   
    print('''[ 0 ] Volks Wagen
    [ 1 ] Hyundai''')

def marca_med():
    print('''[ 0 ] Renault
        [ 1 ] Honda
        [ 2 ] Toyota''')

def versao_modeRena():
    print('''[  ]
    [  ]''')