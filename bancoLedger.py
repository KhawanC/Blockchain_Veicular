import sys
import json
from random import gauss

from numpy import float64
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"


if __name__ == "__main__":
    # Salvar json de dados veicular em uma variavel de nome arq_json
    with open("dadosVeicularesBase.json") as f:
        arq_json = json.load(f)

    # Atualizar todas as emissões dentro do arq_json com base no cálculo de distribuição normal
    for i in arq_json["Modelo_Veiculos"]:
        arq_json["Modelo_Veiculos"][i]["Emissao"] = gauss(
            arq_json["Modelo_Veiculos"][i]["Emissao"], 10)

    print("Dados de emissão atualizados")

    # Criar um novo json com os dados atualizados
    with open('dadosVeicularesAtualizados.json', 'w', ) as arq:
        arq.write(json.dumps(arq_json))

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    c_hlf.new_channel(channel_name)

    for idVeiculo in arq_json["Modelo_Veiculos"]:
        categoria = arq_json["Modelo_Veiculos"][idVeiculo]["Categoria"]
        marca = arq_json["Modelo_Veiculos"][idVeiculo]["Marca"]
        versao = arq_json["Modelo_Veiculos"][idVeiculo]["Versao"]
        modelo = arq_json["Modelo_Veiculos"][idVeiculo]["Modelo"]
        emissao = arq_json["Modelo_Veiculos"][idVeiculo]["Emissao"]
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
            requestor=admin,
            channel_name=channel_name,
            peers=[callpeer],
            cc_name=cc_name,
            cc_version=cc_version,
            fcn='registrarBanco',
            args=[idVeiculo, categoria, marca,
                  versao, modelo, str(emissao)],
            cc_pattern=None))

    print("Successo em registrar seu banco de dados!")
