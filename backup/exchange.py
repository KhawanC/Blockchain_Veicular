import json, random
from scipy.stats import invgauss
from os import path
from IPython.display import display

arq_dir = 'C:/Users/9S/Desktop/testes_1/veiculo.json'
lst = []
tabela = pd.read_excel("C:\\Users\\9S\\Desktop\\testes_1\\lista_carros.xlsx")

if path.isfile(arq_dir) is False:
    raise Exception('Arquivo json faltando')

with open(arq_dir) as fp:
    lst = json.load(fp)

leitor_placa = input('Insira sua placa: ')

if leitor_placa == lst["placa"]:
    print('Acessando veículo...')
    percurso = float(input('Insira a quilometragem do percurso: '))
    emissao = float(input('Insira a média de emissão do veículo: '))
    if percurso == percurso and emissao == float(emissao): 
        print('Calculando crédito...')
        desv_pad = 24,451
        meta = (random.gauss(random.randint(90, 120), 24.451))*percurso
        credito = (emissao*percurso - meta)/1000
        print('Você obteve {:.2f} Créditos de Carbono em sua carteira para uma emissão de {} gCO2 e um percurso de {} Km'.format(credito,emissao,percurso))
        continuar = True
    else:
        print('Os valores devem ser numéricos reais')
        print(percurso, emissao)
        continuar = False

if continuar == True:
    tabela.loc[0, "Placa"] = lst["placa"]
    tabela.loc[0, "Quilometragem Veic. (km)"] = percurso
    tabela.loc[0, "Consumo Veic. (L)"] = lst["placa"]
    tabela.to_excel("lista_carros.xlsx")
    lst.update({
        'creditos': str(credito)
        })
    with open(arq_dir, 'w') as json_file:
        json.dump(lst, json_file,
                        indent=4,
                        separators=(',',': '))
    print('Carteira atualizada')
    display(tabela)
