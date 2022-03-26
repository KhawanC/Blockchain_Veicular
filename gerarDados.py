from random import gauss
from time import sleep
import numpy as np, json
import math, rsa

class Criptografia(object):
    ''' FUNÇÕES DE CRIPTOGRAFIA RSA
    
As próximas linhas são destinadas a criar um par de chaves para criptogravar o arquivo json ao final do código'''
    def generate_keys():
        (pubKey, privKey) = rsa.newkeys(1024)
        with open('keys/pubkey.pem', 'wb') as f:
            f.write(pubKey.save_pkcs1('PEM'))

        with open('keys/privkey.pem', 'wb') as f:
            f.write(privKey.save_pkcs1('PEM'))

    def load_keys():
        with open('keys/pubkey.pem', 'rb') as f:
            pubKey = rsa.PublicKey.load_pkcs1(f.read())

        with open('keys/privkey.pem', 'rb') as f:
            privKey = rsa.PrivateKey.load_pkcs1(f.read())

        return pubKey, privKey

    def encrypt(msg, key):
        return rsa.encrypt(msg.encode('ascii'), key)

    def decrypt(ciphertext, key):
        try:
            return rsa.decrypt(ciphertext, key).decode('ascii')
        except:
            return False

    def sign_sha1(msg, key):
        return rsa.sign(msg.encode('ascii'), key, 'SHA-1')

    def verify_sha1(msg, signature, key):
        try:
            return rsa.verify(msg.encode('ascii'), signature, key) == 'SHA-1'
        except:
            return False
class Json:
    def ler_json(arq_json):
        with open(arq_json, 'r', encoding='utf-8') as f:
            return json.load(f)
    def criar_json(msg):
        with open('TOKEN.json', 'wb') as f:
            f.write(msg)

veiculos_json = Json.ler_json('dadosVeiculares.json')
sigma = 10
qst = ["categoria", "marca"] 
user_temp = []
user_fnl = {}
chave = list(veiculos_json.keys())

#Ler placa e verificr se a placa foi digitada corretamente
placa = input('Insira sua placa: ')
sleep(1)
print("-="*5)
opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
while int(opt) == 0:
    sleep(1)
    print('-='*5)
    placa = input('Insira sua placa: ')
    opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')
    while int(opt) < 0 or int(opt) > 1:
        sleep(1)
        print('-='*5)
        print('O valor deve ser entre 1 e 2')
        opt = input('Tem certeza que essa é sua placa?\n[ 0 ] Não\n[ 1 ] Sim\n')

#Ler categoria do veiculo
for i in range(0, 2):
    print("-="*5)
    print("\nQual a/o {} do seu veículo?\n".format(qst[i]))
    sleep(1)
    if i == 0:
        user_temp.append(int(input('''[ 0 ] Compacto
[ 1 ] Médio\n: ''')))
    #Ler marca do veículo caso e indicar corretamente caso seja da categoria médio ou compacto
    if i == 1 and user_temp[0] == 0:
        sleep(1)
        for j in range(0,3):
            print("[ {} ] {}".format(j, veiculos_json["Veiculo_Compacto"][j]["Marca"]))
        user_temp.append(int(input("\n: " )))
    if i == 1 and user_temp[0] == 1:
        sleep(1)
        for j in range(0,3):
            print("[ {} ] {}".format(j, veiculos_json["Veiculo_Medio"][j]["Marca"]))
        user_temp.append(int(input("\n: " )))

user_fnl["Placa"] = placa
user_fnl["Categoria"] = chave[user_temp[0]]
user_fnl = user_fnl | veiculos_json["{}".format(user_fnl["Categoria"])][user_temp[1]]

Criptografia.generate_keys()
Criptografia.pubKey, Criptografia.privKey = Criptografia.load_keys()
Json.criar_json(Criptografia.encrypt(str(user_fnl), Criptografia.pubKey))