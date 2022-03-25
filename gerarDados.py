from random import gauss
import numpy as np, json
import math

def ler_json(arq_json):
    with open(arq_json, 'r', encoding='utf-8') as f:
        return json.load(f)

veiculos_json = ler_json('dadosVeiculares.json')
sigma = 10
qst = ["categoria", "marca"] 
user_temp = []
user_fnl = {}
chave = list(veiculos_json.keys())

#Ler placa e verificr se a placa foi digitada corretamente
placa = input('Insira sua placa: ')
print("-="*5)
opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
while int(opt) == 0:
    print('-='*5)
    placa = input('Insira sua placa: ')
    opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
    while int(opt) < 0 or int(opt) > 1:
        print('-='*5)
        print('O valor deve ser entre 1 e 2')
        opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')

#Ler categoria do veiculo
for i in range(0, 2):
    print("-="*5)
    print("\nQual o/a {} do seu veículo?\n".format(qst[i]))
    if i == 0:
        user_temp.append(int(input('''[ 0 ] Compacto
[ 1 ] Médio\n: ''')))
    #Ler marca do veículo caso e indicar corretamente caso seja da categoria médio ou compacto
    if i == 1 and user_temp[0] == 0:
        for j in range(0,3):
            print("[ {} ] {}".format(j, veiculos_json["Veiculo_Compacto"][j]["Marca"]))
        user_temp.append(int(input("\n: " )))
    if i == 1 and user_temp[0] == 1:
        for j in range(0,3):
            print("[ {} ] {}".format(j, veiculos_json["Veiculo_Medio"][j]["Marca"]))
        user_temp.append(int(input("\n: " )))

user_fnl["Placa"] = placa
user_fnl["Categoria"] = chave[user_temp[0]]
user_fnl = user_fnl | veiculos_json["{}".format(user_fnl["Categoria"])][user_temp[1]]