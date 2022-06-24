import couchdb, asyncio
from hfc.fabric import Client as client_fabric

domain = "ptb.de"
channel_name = "nmi-channel"
cc_name = "fabpki"
cc_version = "1.0"
couch = couchdb.Server()
db = couch['nmi-channel_fabpki']
listaFabricantes = []
listaOrdemVenda = []

if __name__ == "__main__":
    
    loop = asyncio.get_event_loop()
    c_hlf = client_fabric(net_profile=(domain + ".json"))
    admin = c_hlf.get_user(domain, 'Admin')
    callpeer = "peer0." + domain
    c_hlf.new_channel(channel_name)
    
    for doc in db.view('_all_docs'):
        i = doc['id']
        if i[0:4] == "fab-":
            listaFabricantes.append(i[4:])
            
    nomeFabricante = input("Qual fabricante deseja criar a ordem? ")
    if nomeFabricante.upper().replace(" ", "-") in listaFabricantes:
        tipoTransacao = input("Você deseja comprar ou vender carbono? ")
        if tipoTransacao.lower() == "vender":
            quantidade = input("Quanto de carbono você deseja vender? ")
            if float(quantidade) > 0:
                print("processando...")
                listaOrdemVenda.append(nomeFabricante.upper().replace(" ", "-"))
                listaOrdemVenda.append(tipoTransacao.lower())
                listaOrdemVenda.append(float(quantidade))
            else:
                print("Valor inválido")
                exit()
        if tipoTransacao.lower() == "comprar":
            quantidade = input("Qual a sua oferta para comprar carbono? ")
            if float(quantidade) > 0:
                print("processando...")
                listaOrdemVenda.append(nomeFabricante.upper().replace(" ", "-"))
                listaOrdemVenda.append(tipoTransacao.lower())
                listaOrdemVenda.append(float(quantidade))
            else:
                print("Valor inválido")
                exit()
        if tipoTransacao.lower() != "comprar" and tipoTransacao.lower() != "vender":
            print("O tipo da ordem deve ser de compra ou venda")
            exit()
    else:
        print("Não achamos esse fabricante")
        print(listaFabricantes)
        exit()
        
    if len(listaOrdemVenda) == 3:
        print(listaOrdemVenda)
        
        response = loop.run_until_complete(c_hlf.chaincode_invoke(
                requestor=admin,
                channel_name=channel_name,
                peers=[callpeer],
                cc_name=cc_name,
                cc_version=cc_version,
                fcn='anunciarOrdem',
                args=[listaOrdemVenda[0], listaOrdemVenda[1], str(listaOrdemVenda[2])],
                cc_pattern=None))
        
        print("Ordem anunciada com sucesso!")