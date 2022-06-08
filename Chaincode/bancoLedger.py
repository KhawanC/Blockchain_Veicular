import json, asyncio, secrets
from random import gauss
from hfc.fabric import Client as client_fabric

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
montadoras = []

if __name__ == "__main__":
    # Salvar json de dados veicular em uma variavel de nome arq_json
    with open("dadosVeicularesBase.json", 'r', encoding='utf8') as f:
        arq_json = json.load(f)

    # Atualizar todas as emissões dentro do arq_json com base no cálculo de distribuição normal    
    for i in range(len(arq_json)):
        arq_json[i]["G.E._Gasolina_Diesel_CO2_fossil_(g/km)"] = gauss(
            arq_json[i]["G.E._Gasolina_Diesel_CO2_fossil_(g/km)"], 10)

    print("Dados de emissão atualizados")

    # Criar um novo json com os dados atualizados
    with open('dadosVeicularesAtualizados.json', 'w', ) as arq:
        json.dump(arq_json, arq, indent=2, separators=(',', ': '), ensure_ascii=False)

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    for indexVeiculo in range(len(arq_json)):        
        a = arq_json[indexVeiculo]["Categoria"]
        b = arq_json[indexVeiculo]["Marca"]
        c = arq_json[indexVeiculo]["Modelo"]
        d = arq_json[indexVeiculo]["Versao"]
        e = arq_json[indexVeiculo]["Motor"]
        f = arq_json[indexVeiculo]["Tipo_de_Propulsao"]
        g = arq_json[indexVeiculo]["Transmissao_Velocidades_(n)"]
        h = arq_json[indexVeiculo]["Ar_Condicionado"]
        i = arq_json[indexVeiculo]["Direcao_Assistida"]
        j = arq_json[indexVeiculo]["Combustivel"]
        k = arq_json[indexVeiculo]["NMOG_Nox_(mg/km)"]
        l = arq_json[indexVeiculo]["CO_(mg/km)"]
        m = arq_json[indexVeiculo]["CHO_(mg/km)"]
        n = arq_json[indexVeiculo]["Reducao_Relativa_Ao_Limite"]
        o = arq_json[indexVeiculo]["G.E._Etanol_CO2_Fossil_(g/km)"]
        p = arq_json[indexVeiculo]["G.E._Gasolina_Diesel_CO2_fossil_(g/km)"]
        if f == "Eletrico" or p == 0: #Verificar se o carro é Elétrico, caso seja, ignore
            continue
        q = arq_json[indexVeiculo]["Etanol_-_Cidade_(km/l)"]
        r = arq_json[indexVeiculo]["Etanol_-_Estrada_(km/l)"]
        s = arq_json[indexVeiculo]["Gasolina_Diesel_Eletrico_-_Cidade_(km/l)"]
        t = arq_json[indexVeiculo]["Gasolina_Diesel_Eletrico_-_Estrada_(km/l)"]
        u = arq_json[indexVeiculo]["Consumo_Energetico_(MJ/km)"]
        v = arq_json[indexVeiculo]["Classificacao_PBE_-_Relativo_na_Categoria"]
        w = arq_json[indexVeiculo]["Classificação_PBE_-_Absoluto_Geral"]
        x = arq_json[indexVeiculo]["Selo_CONPET_de_Eficiencia_Energetica"]
        
        carroExiste = False
        
        for i in montadoras: #Loop para inserir o nome da montadora, caso esse nome estaja nao esteja lista de montadora é ecessário inserir, caso o contrário não é necessário
            if i == b:
                carroExiste = True
        if carroExiste == False:
            montadoras.append(b)
        
        cdg = a[0:3] + b + str(c) + str(e) + f[0:3] + str(g) + str(h) + j + secrets.token_urlsafe(16)

        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarModelo',
            args=[cdg.replace(" ", "-"),str(a),str(b),str(c),str(d),str(p)],
            cc_pattern=None))

    for i in montadoras:
        loop = asyncio.get_event_loop()

        c_hlf = client_fabric(net_profile=(domain + ".json"))

        admin = c_hlf.get_user(domain, 'Admin')
        callpeer = "peer0." + domain

        c_hlf.new_channel(channel_name)
        
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarFabricante',
            args=[i],
            cc_pattern=None))
        
    print("Successo em registrar seu banco de dados!")
