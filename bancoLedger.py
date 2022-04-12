import sys
import json
from random import gauss
from hfc.fabric import Client as client_fabric
import asyncio

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"


if __name__ == "__main__":

    arq_json = {"Veiculo":
                [{
                    "Categoria": "Compacto",
                    "Marca": "VW",
                    "Versao": "GOL",
                    "Modelo": "PATRULHEIRO",
                    "Emissao": 112,
                    "Codigo": "CVKWGOLPAT112"
                },
                    {
                    "Categoria": "Compacto",
                    "Marca": "Hyundai",
                    "Versao": "HB20",
                    "Modelo": "Vision21/22",
                    "Emissao": 97,
                    "Codigo": "CHYDHB2VIS97"
                },
                    {
                    "Categoria": "Compacto",
                    "Marca": "Toyota",
                    "Versao": "Etios Hatback",
                    "Modelo": "XSTD",
                    "Emissao": 100,
                    "Codigo": "CTYTETHXST100"
                },
                    {
                    "Categoria": "Compacto",
                    "Marca": "Ford",
                    "Versao": "Ka Hatch",
                    "Modelo": "SE",
                    "Emissao": 92,
                    "Codigo": "CFRDKAHSEE92"
                },
                    {
                    "Categoria": "Compacto",
                    "Marca": "MINI",
                    "Versao": "Cooper",
                    "Modelo": "5P",
                    "Emissao": 105,
                    "Codigo": "CMNICPR5PP105"
                },
                    {
                    "Categoria": "Compacto",
                    "Marca": "Renault",
                    "Versao": "Sandero",
                    "Modelo": "Life",
                    "Emissao": 97,
                    "Codigo": "CRNTSDRLFE97"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "Renault",
                    "Versao": "Logan",
                    "Modelo": "Zen",
                    "Emissao": 107,
                    "Codigo": "MRNTLOGZEN107"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "Honda",
                    "Versao": "City",
                    "Modelo": "EX",
                    "Emissao": 101,
                    "Codigo": "MHNDCTYEXX101"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "Toyota",
                    "Versao": "Etius Seda",
                    "Modelo": "XSTD",
                    "Emissao": 98,
                    "Codigo": "MTYTETSXST98"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "NIISAN",
                    "Versao": "V-DRIVE",
                    "Modelo": "16 MT",
                    "Emissao": 99,
                    "Codigo": "MNSNVDR16M99"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "KIA",
                    "Versao": "RIO",
                    "Modelo": "EX 1.6 FF AT HB",
                    "Emissao": 115,
                    "Codigo": "MKIARIOEXF115"
                },
                    {
                    "Categoria": "Medio",
                    "Marca": "CHEVROLET",
                    "Versao": "ONIX PLUS",
                    "Modelo": "3LT",
                    "Emissao": 88,
                    "Codigo": "MCVLOXP3LT88"
                }]}
    for i in range(0, 12):
        arq_json["Veiculo"][i]["Emissao"] = gauss(
            arq_json["Veiculo"][i]["Emissao"], 10)

    print("Dados de emissão atualizados")
    with open('dadosVeiculares.json', 'w', ) as arq:
        arq.write(json.dumps(arq_json))

    arq_json = str(arq_json)

    loop = asyncio.get_event_loop()

    c_hlf = client_fabric(net_profile=(domain + ".json"))

    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain

    print("Checando instalação de arquivo fbpki:")
    response = loop.run_until_complete(c_hlf.query_installed_chaincodes(
        requestor=admin,
        peers=[callpeer]
    ))
    print(response)

    c_hlf.new_channel(channel_name)

    response = loop.run_until_complete(c_hlf.chaincode_invoke(
        requestor=admin,
        channel_name=channel_name,
        peers=[callpeer],
        cc_name=cc_name,
        cc_version=cc_version,
        fcn='registrarBanco',
        args=[arq_json],
        cc_pattern=None))

    print("Successo em registrar seu banco de dados!")
