/////////////////////////////////////////////
//    THE BLOCKCHAIN PKI EXPERIMENT     ////
///////////////////////////////////////////
/*
	This is the fabpki, a chaincode that implements a Public Key Infrastructure (PKI)
	for measuring instruments. It runs in Hyperledger Fabric 1.4.
	He was created as part of the PKI Experiment. You can invoke its methods
	to store measuring instruments public keys in the ledger, and also to verify
	digital signatures that are supposed to come from these instruments.

	@author: Wilson S. Melo Jr.
	@date: Oct/2019
*/
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890") //Criar sequencia de letras

func AleatString(n int) string { //Função para criar uma sequencia de números aleatórios
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Encode(msg string) string { //Função para criar um hash sha-1
	h := sha1.New()
	h.Write([]byte(msg))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func Arredondar(n float64) float64 {
	numeroArredondado := float64(int(n*10000)) / 10000
	return numeroArredondado
}

type SmartContract struct {
}

type ModeloPBE struct { //"model-"
	Categoria  string  `json:"Categoria"`
	Fabricante string  `json:"Fabricante"`
	Versao     string  `json:"Versao"`
	Modelo     string  `json:"Modelo"`
	Emissao    float64 `json:"Emissao"`
}

type VeiculoPBE struct { //"veic-"
	IdModelo   string  `json:"IdModelo"`
	Co2Emitido float64 `json:"Co2Emitido"`
}

type trajetoPBE struct { //traj-
	Distancia float64 `json:"Distancia"`
	idVeiculo string  `json:"idVeiculo"`
}

type Veiculo struct { //"veic-"
	Hash       string  `json:"Hash"`
	Co2Emitido float64 `json:"Co2Emitido"`
	Fabricante string  `json:"Fabricante"`
}

type Fabricante struct { //"fab-""
	Co2Tot          float64 `json:"Co2_Tot"`
	SaldoCarbono    float64 `json:"SaldoCarbono"`
	SaldoFiduciario float64 `json:"Saldo_Fiduciario"`
}

type OrdemTransacao struct { //"trans-"
	ProprietarioOrdem string  `json:"ProprietarioOrdem"` // FK (Veiculo)
	TipoTransacao     string  `json:"TipoTransacao"`     // 1: Vender carbono -- 2: Comprar carbono
	SaldoOfertado     float64 `json:"SaldoOfertado"`
	IdComprador       string  `json:"IdComprador"`
	ValorUltimoLance  float64 `json:"ValorUltimoLance"`
	StatusOrdem       string  `json:"StatusOrdem"` // Recente - Andamento - Fechado
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registrarVeiculo" {
		return s.registrarVeiculo(stub, args)
	} else if fn == "registrarVeiculoPBE" {
		return s.registrarVeiculoPBE(stub, args)
	} else if fn == "registrarModeloPBE" {
		return s.registrarModeloPBE(stub, args)
	} else if fn == "registrarFabricante" {
		return s.registrarFabricante(stub, args)
	} else if fn == "registrarTrajetoPBE" {
		return s.registrarTrajetoPBE(stub, args)
	} else if fn == "registrarCredito" {
		return s.registrarCredito(stub, args)
	} else if fn == "anunciarOrdem" {
		return s.anunciarOrdem(stub, args)
	} else if fn == "ordemLance" {
		return s.ordemLance(stub, args)
	} else if fn == "fecharOrdem" {
		return s.fecharOrdem(stub, args)
	}

	return shim.Error("Chaincode não suporta essa função.")
}

func (s *SmartContract) registrarFabricante(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Era esperado 1 único argumento... Tente novamente!")
	}

	nomeFab := args[0]

	fabricanteInfor := Fabricante{
		Co2Tot:          0.0,
		SaldoCarbono:    0.0,
		SaldoFiduciario: 100000.0,
	}

	fabricanteAsBytes, _ := json.Marshal(fabricanteInfor)

	idCdgLedger := "fab-" + nomeFab

	stub.PutState(idCdgLedger, fabricanteAsBytes)

	fmt.Println("Sucesso ao registrar fabricantes")
	return shim.Success(nil)
}

func (s *SmartContract) registrarVeiculo(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Eram esperados 4 argumentos... Tente novamente!")
	}

	vim := args[0]
	hash := args[1]
	co2 := args[2]
	fabNome := args[3]

	co2VeicFloat, err := strconv.ParseFloat(co2, 64)

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Veiculo{
		Hash:       hash,
		Co2Emitido: co2VeicFloat,
		Fabricante: fabNome,
	}

	//Recuperando dados do usuário
	fabricanteAsBytes, err := stub.GetState(("fab-" + fabNome))
	if err != nil || fabricanteAsBytes == nil {

		fabricanteInfor := Fabricante{
			Co2Tot:          0.0,
			SaldoCarbono:    0.0,
			SaldoFiduciario: 100000.0,
		}

		fabricanteInfor.Co2Tot += co2VeicFloat

		veiculoAsBytes, _ := json.Marshal(userVeiculo)
		fabricanteAsBytes, _ = json.Marshal(fabricanteInfor)

		stub.PutState(("fab-" + fabNome), fabricanteAsBytes)
		stub.PutState(("veic-" + vim), veiculoAsBytes)

		fmt.Println("Sucesso ao registrar veiculo e fabricante")
		return shim.Success(nil)

	}

	//Criando Struct para encapsular os dados do veiculo
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	fabricante.Co2Tot += co2VeicFloat

	veiculoAsBytes, _ := json.Marshal(userVeiculo)
	fabricanteAsBytes, _ = json.Marshal(fabricante)

	stub.PutState(("fab-" + fabNome), fabricanteAsBytes)
	stub.PutState(("veic-" + vim), veiculoAsBytes)

	fmt.Println("Sucesso ao registrar veiculo e fabricante")
	return shim.Success(nil)
}

func (s *SmartContract) registrarVeiculoPBE(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Eram esperado 2 argumentos... Tente novamente!")
	}

	placa := args[0]
	idModelo := args[1]

	userVeiculo := VeiculoPBE{
		IdModelo:   idModelo,
		Co2Emitido: 0,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)
	stub.PutState(("veic-" + placa), veiculoAsBytes)

	fmt.Println("Sucesso ao registrar veiculo")
	return shim.Success(nil)
}

func (s *SmartContract) registrarModeloPBE(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Eram esperado 6 argumentos... Tente novamente!")
	}

	idModelo := args[0]
	categoria := args[1]
	fabricante := args[2]
	versao := args[3]
	modelo := args[4]
	co2Emissao := args[5]

	co2ModelFloat, err := strconv.ParseFloat(co2Emissao, 64)

	if err != nil {
		shim.Error("Não foi possível converter a quantidade de co2")
	}

	modeloInfor := ModeloPBE{
		Categoria:  categoria,
		Fabricante: fabricante,
		Versao:     versao,
		Modelo:     modelo,
		Emissao:    co2ModelFloat,
	}

	modeloAsBytes, _ := json.Marshal(modeloInfor)

	stub.PutState(("model-" + idModelo), modeloAsBytes)

	fmt.Println("Modelo registrado com sucesso")
	return shim.Success(nil)
}

func (s *SmartContract) registrarTrajetoPBE(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Eram esperado 3 argumentos... Tente novamente!")
	}

	distancia := args[0]
	idVeiculoPBE := args[1]
	idTrajeto := "traj-" + Encode(AleatString(10))

	distanciaFloat, err := strconv.ParseFloat(distancia, 64)

	veiculoPBEAsBytes, err := stub.GetState(idVeiculoPBE)
	if err != nil || veiculoPBEAsBytes == nil {
		return shim.Error("Não foi possível resgatar a placa do seu veículo")
	}

	veiculo := VeiculoPBE{}
	json.Unmarshal(veiculoPBEAsBytes, &veiculo)

	modeloPBEAsBytes, err := stub.GetState(("model-" + veiculo.IdModelo))
	if err != nil || modeloPBEAsBytes == nil {
		return shim.Error("Não foi possível resgatar o modelo do seu veículo")
	}

	modelo := ModeloPBE{}
	json.Unmarshal(modeloPBEAsBytes, &modelo)

	fabricanteAsBytes, err := stub.GetState(("fab-" + modelo.Fabricante))
	if err != nil || fabricanteAsBytes == nil {
		return shim.Error("Não foi possível resgatar o fabricante do seu veículo")
	}

	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	trajeto := trajetoPBE{
		idVeiculo: idVeiculoPBE,
		Distancia: distanciaFloat,
	}

	carbonoEmitido := trajeto.Distancia * modelo.Emissao
	veiculo.Co2Emitido += carbonoEmitido
	fabricante.Co2Tot += carbonoEmitido

	veiculoPBEAsBytes, _ = json.Marshal(veiculo)
	fabricanteAsBytes, _ = json.Marshal(fabricante)
	trajetoAsBytes, _ := json.Marshal(trajeto)

	stub.PutState(idVeiculoPBE, veiculoPBEAsBytes)
	stub.PutState(("fab-" + modelo.Fabricante), fabricanteAsBytes)
	stub.PutState(idTrajeto, trajetoAsBytes)

	fmt.Println("Sucesso ao registrar trajeto e emissão de co2")
	return shim.Success(nil)
}

func (s *SmartContract) registrarCredito(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Era esperado 1 único argumento... Tente novamente!")
	}

	idFabricante := args[0]

	//Recuperando dados do usuário
	fabricanteAsBytes, err := stub.GetState(idFabricante)
	if err != nil || fabricanteAsBytes == nil {
		return shim.Error("Seu fabricante não existe.")
	}

	//Criando Struct para encapsular os dados do veiculo
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	if fabricante.Co2Tot == 0.0 {
		fmt.Println("Não foi computado emissão de carbono para o fabricante: " + idFabricante)
		return shim.Success(nil)
	}

	var saldo = 50000.0 - fabricante.Co2Tot
	fabricante.SaldoCarbono = saldo
	fabricante.Co2Tot = 0.0

	fabricanteAsBytes, _ = json.Marshal(fabricante)

	stub.PutState(idFabricante, fabricanteAsBytes)

	fmt.Println("Saldo de carbono computado com sucesso: " + idFabricante)
	return shim.Success(nil)
}

func (s *SmartContract) anunciarOrdem(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		fmt.Println("Eram esperados 3 arugmentos")
		return shim.Error("Era esperado 3 argumentos... Tente novamente!")
	}

	proprietarioOrdem := args[0]
	tipoTransacao := args[1]
	saldoOferta := args[2]

	saldoOfertaFloat, err := strconv.ParseFloat(saldoOferta, 64)

	//Verificando se o fabricante realmente existe
	fabricanteAsBytes, err := stub.GetState(proprietarioOrdem)
	if err != nil || fabricanteAsBytes == nil {
		fmt.Println("Seu fabricante não existe")
		return shim.Error("Seu fabricante não existe.")
	}

	//Criando Struct para encapsular os dados do fabricante
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	if tipoTransacao == "vender" {
		if saldoOfertaFloat > fabricante.SaldoCarbono {
			fmt.Println("Você não tem saldo de carbono suficiente")
			return shim.Error("Você não tem saldo de carbono suficiente")
		}
		fabricante.SaldoCarbono -= saldoOfertaFloat
	}

	if tipoTransacao == "comprar" {
		if saldoOfertaFloat > fabricante.SaldoFiduciario {
			fmt.Println("VocÊ não tem saldo fiduciario suficiente")
			return shim.Error("Você não tem saldo fiduciario suficiente")
		}
		fabricante.SaldoFiduciario -= saldoOfertaFloat
	}

	ordemVenda := OrdemTransacao{
		IdComprador:       "null",
		ValorUltimoLance:  0.0,
		StatusOrdem:       "Recente",
		ProprietarioOrdem: proprietarioOrdem,
		SaldoOfertado:     saldoOfertaFloat,
		TipoTransacao:     tipoTransacao,
	}

	ordemVendaAsBytes, _ := json.Marshal(ordemVenda)
	fabricanteAsBytes, _ = json.Marshal(fabricante)

	idOrdem := "trans-" + Encode(AleatString(10))

	stub.PutState(proprietarioOrdem, fabricanteAsBytes)
	stub.PutState(idOrdem, ordemVendaAsBytes)

	fmt.Println("Ordem de " + tipoTransacao + " anunciado com sucesso!")
	return shim.Success(nil)
}

func (s *SmartContract) ordemLance(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		fmt.Println("Eram esperados 2 argumentos")
		return shim.Error("Era esperado 2 argumentos... Tente novamente!")
	}

	idTransacao := args[0]
	valorLance := args[1]
	idComprador := args[2]

	valorLanceFloat, err := strconv.ParseFloat(valorLance, 64)

	//Recuperando dados da transação
	ordemTransacaoAsBytes, err := stub.GetState(idTransacao)
	if err != nil || ordemTransacaoAsBytes == nil {
		fmt.Println("Seu proprietario nao existe")
		return shim.Error("Seu proprietário não existe.")
	}

	//Recuperando dados do proprietário
	fabricanteAsBytes, err := stub.GetState(idComprador)
	if err != nil || ordemTransacaoAsBytes == nil {
		fmt.Println("Seu fabricante nao existe")
		return shim.Error("Seu fabricante não existe.")
	}

	//Encapsulando os dados da ordem de transação e do fabricante
	ordem := OrdemTransacao{}
	json.Unmarshal(ordemTransacaoAsBytes, &ordem)

	if ordem.StatusOrdem == "Fechado" {
		fmt.Println("Essa ordem está fechada")
		return shim.Error("Essa ordem não pode mais ser movimentado pois o proprietário à fechou.")
	}

	if idComprador == ordem.IdComprador {
		fmt.Println("Você não pode efetuar um lance duas vezes")
	}

	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	if valorLanceFloat > fabricante.SaldoFiduciario && ordem.TipoTransacao == "vender" {
		fmt.Println("Voce nao tem saldo fiduciario suficiente")
		return shim.Error("Você não tem saldo fiduciario suficiente.")
	}

	if valorLanceFloat > fabricante.SaldoCarbono && ordem.TipoTransacao == "comprar" {
		fmt.Println("Você nao tem saldo de carbono suficiente")
		return shim.Error("Você não tem saldo de carbono suficiente.")
	}

	if ordem.ValorUltimoLance > valorLanceFloat {
		fmt.Println("Seu lance é menor do que o lance anterior")
		return shim.Error("Seu lance é menor do que o lance anterior.")
	}

	ordem.StatusOrdem = "Andamento"
	ordem.ValorUltimoLance = valorLanceFloat
	ordem.IdComprador = idComprador

	ordemTransacaoAsBytes, _ = json.Marshal(ordem)
	stub.PutState(idTransacao, ordemTransacaoAsBytes)

	fmt.Println("Lance registrado no sucesso")

	return shim.Success(nil)
}

func (s *SmartContract) fecharOrdem(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		fmt.Println("Eram esperados 2 argumentos")
		return shim.Error("Era esperado 2 argumentos... Tente novamente!")
	}

	idTransacao := args[0]
	idProprietario := args[1]

	//Recuperando dados da transação
	ordemTransacaoAsBytes, err := stub.GetState(idTransacao)
	if err != nil || ordemTransacaoAsBytes == nil {
		fmt.Println("Essa ordem nao existe")
		return shim.Error("Essa ordem não existe.")
	}

	//Encapsulando os dados do fabricante
	ordem := OrdemTransacao{}
	json.Unmarshal(ordemTransacaoAsBytes, &ordem)

	if ordem.ProprietarioOrdem != idProprietario {
		fmt.Println("Você nao é proprietario dessa ordem")
		return shim.Error("Você não é o proprietário dessa ordem")
	}

	//Recuperando dados do proprietario
	proprietarioAsBytes, err := stub.GetState(ordem.ProprietarioOrdem)
	if err != nil || ordemTransacaoAsBytes == nil {
		fmt.Println("Seu proprietario nao existe")
		return shim.Error("Seu proprietário não existe.")
	}

	//Encapsulando os dados do fabricante
	proprietario := Fabricante{}
	json.Unmarshal(proprietarioAsBytes, &proprietario)

	//Recuperando dados do comprador
	compradorAsBytes, err := stub.GetState(ordem.IdComprador)
	if err != nil || ordemTransacaoAsBytes == nil {
		fmt.Println("Seu comprador nao existe")
		return shim.Error("Seu comprador não existe.")
	}

	//Encapsulando os dados do fabricante
	comprador := Fabricante{}
	json.Unmarshal(compradorAsBytes, &comprador)

	if ordem.IdComprador == "null" {
		fmt.Println("Nao houveram lances para essa ordem")

		return shim.Error("Não houveram lances para essa ordem")
	}

	if ordem.TipoTransacao == "vender" {
		proprietario.SaldoFiduciario += ordem.ValorUltimoLance
		comprador.SaldoFiduciario -= ordem.ValorUltimoLance
		ordem.StatusOrdem = "Fechado"
	}

	if ordem.TipoTransacao == "comprar" {
		proprietario.SaldoCarbono += ordem.ValorUltimoLance
		comprador.SaldoCarbono -= ordem.ValorUltimoLance
		ordem.StatusOrdem = "Fechado"
	}

	ordemTransacaoAsBytes, _ = json.Marshal(ordem)
	proprietarioAsBytes, _ = json.Marshal(proprietario)
	compradorAsBytes, _ = json.Marshal(comprador)

	stub.PutState(idTransacao, ordemTransacaoAsBytes)
	stub.PutState(idProprietario, proprietarioAsBytes)
	stub.PutState(ordem.IdComprador, compradorAsBytes)

	fmt.Println("Transação finalizada com sucesso")
	return shim.Success(nil)

}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
