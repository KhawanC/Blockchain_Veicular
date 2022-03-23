from random import gauss
import numpy as np

import math, openpyxl as exl

tabela = exl.load_workbook(filename= 'lista_carros.xlsx')
tabela_act = tabela.active

sigma = 10

for cel in tabela_act['F']:
    try:
        aleat_cel = gauss(float(cel.value),sigma)
        linha = cel.row
        tabela_act[f'F{linha}'] = aleat_cel
    except:
        print('error')

tabela.save('lista_carros(2).xlsx')

