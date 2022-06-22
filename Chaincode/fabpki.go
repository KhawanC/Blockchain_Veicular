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

type ModeloVeiculo struct { //"model-"
	Categoria  string `json:"Categoria"`
	Marca      string `json:"Marca"`
	Versao     string `json:"Versao"`
	Modelo     string `json:"Modelo"`
	EmissaoCo2 string `json:"EmissaoCo2"`
}

type Veiculo struct { //"user-"
	IdCdgModelo         string `json:"IdCdgModelo"` //FK (Categoria)
	AcumuladorDistancia string `json:"AcumuladorDistancia"`
}

type Veiculo2 struct { //"user-"
	Hash string `json:"Hash"`
	Vim  string `json:"Vim"`
	Co2  string `json:"Co2"`
}

type Trajeto struct { //"traj-"
	TrajetoDistancia string `json:"TrajetoDistancia"`
}

type TrajetoVeiculo struct { //"traj_user-"
	IdPlaca   string `json:"IdPlaca"`   //FK (Veiculo)
	IdTrajeto string `json:"IdTrajeto"` //FK (Trajeto)
}

type Fabricante struct { //"fab-""
	Co2Tot          string `json:"Co2_Tot"`
	SaldoCarbono    string `json:"SaldoCarbono"`
	SaldoFiduciario string `json:"Saldo_FIduciario"`
}

type OrdemTransacao struct { //"trans-"
	ProprietarioOrdem string `json:"ProprietarioOrdem"` // FK (Veiculo)
	TipoTransacao     string `json:"TipoTransacao"`     // 1: Vender carbono -- 2: Comprar carbono
	SaldoOfertado     string `json:"SaldoOfertado"`
	IdComprador       string `json:"IdComprador"`
	ValorLance        string `json:"ValorLance"`
	StatusOrdem       string `json:"StatusOrdem"` // Recente - Andamento - Fechado
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registrarModelo" {
		return s.registrarModelo(stub, args)
	} else if fn == "registrarVeiculo" {
		return s.registrarVeiculo(stub, args)
	} else if fn == "registrarTrajeto" {
		return s.registrarTrajeto(stub, args)
	} else if fn == "registrarFabricante" {
		return s.registrarFabricante(stub, args)
	} else if fn == "registrarCarbono" {
		return s.registrarCarbono(stub, args)
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

//Função que recebe ./bancoLedger.py e inserer o banco de dados com os veículos
func (s *SmartContract) registrarModelo(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificando se a quantidade de argumnetos é maior que 6
	if len(args) != 6 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	//Inserindo argumentos dentro de variáveis
	cdg := args[0]
	categoria := args[1]
	marca := args[2]
	versao := args[3]
	modelo := args[4]
	emissao_co2 := args[5]

	var categoriaInfor = ModeloVeiculo{ //Inserindo argumentos dentro da Struct Categoria
		Categoria:  categoria,
		Marca:      marca,
		Versao:     versao,
		Modelo:     modelo,
		EmissaoCo2: emissao_co2,
	}

	categoriaAsBytes, _ := json.Marshal(categoriaInfor) //Encapsulando as informações em arquivo JSON

	idCdgLedger := "model-" + cdg

	stub.PutState(idCdgLedger, categoriaAsBytes) //Inserindo valores no ledger, com uma informação associada à uma chave

	fmt.Println("Sucesso ao registrar  modelo de veiculo")
	return shim.Success(nil)
}

func (s *SmartContract) registrarFabricante(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	nomeFab := args[0]

	//Verificando se a quantidade de argumnetos é maior que 1
	if len(args) != 1 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	fabricanteInfor := Fabricante{
		Co2Tot:          "0.0",
		SaldoCarbono:    "0.0",
		SaldoFiduciario: "10000.0",
	}

	fabricanteAsBytes, _ := json.Marshal(fabricanteInfor) //Encapsulando as informações em arquivo JSON

	idCdgLedger := "fab-" + nomeFab

	stub.PutState(idCdgLedger, fabricanteAsBytes) //Inserindo valores no ledger, com uma informação associada à uma chave

	fmt.Println("Sucesso ao registrar fabricantes")
	return shim.Success(nil)
}

//Função que recebe "./userLedger" e insere o veículo do usuário no Ledger
func (s *SmartContract) registrarVeiculo(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se existem mais de 2 argumentos no código do cliente
	if len(args) != 3 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}

	vim := args[0]
	hash := args[1]
	co2 := args[2]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Veiculo2{
		Hash: hash,
		Vim:  vim,
		Co2:  co2,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)

	idUserVeiculo := "veic-" + vim
	stub.PutState(idUserVeiculo, veiculoAsBytes)

	fmt.Println("Sucesso ao registrar veiculo")
	return shim.Success(nil)
}

func (s *SmartContract) registrarTrajeto(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se arquivo py retornou 2 argumentos
	if len(args) != 2 {
		return shim.Error("Eram esperados 2 argumentos... Tente novamente!")
	}

	//Dar nome aos argumentos
	idPlaca := args[0]
	userDistancia := args[1]

	cdgUnicoTrajeto := Encode(AleatString(20))                  //Criando código unico para Struct trajeto
	cdgUnicoUsuarioTrajeto := Encode(AleatString(20))           //Criando código unico para Struct trajeto
	userDistFLoat, err := strconv.ParseFloat(userDistancia, 64) //Convertendo a distância recebida para float

	//Recuperando dados do usuário
	userAsBytes, err := stub.GetState(idPlaca)
	if err != nil || userAsBytes == nil {
		return shim.Error("Sua placa não existe.")
	}

	//Criando Struct para encapsular os dados
	veiculo := Veiculo{}
	json.Unmarshal(userAsBytes, &veiculo)

	//Criar Struct do trajeto e do usuário
	trajetoVeiculo := TrajetoVeiculo{}
	trajeto := Trajeto{}

	distAcumuladoString := veiculo.AcumuladorDistancia //Inserindo o acumulador de distância do usuário dentro de uma variável
	distAcumuladoFloat, err := strconv.ParseFloat(distAcumuladoString, 64)

	distAcumuladoFloat += userDistFLoat //Adicionar a distancia do trajeto feito ao acumulador
	newDist := fmt.Sprintf("%g", distAcumuladoFloat)

	veiculo.AcumuladorDistancia = newDist

	//Criar assinatura do trajeto
	trajeto.TrajetoDistancia = userDistancia

	//Associar trajeto com o usuário
	trajetoVeiculo.IdPlaca = idPlaca
	trajetoVeiculo.IdTrajeto = "trajeto-" + cdgUnicoTrajeto

	//Encapsulando dados em arquivo JSON
	veiculoAsBytesFinal, _ := json.Marshal(veiculo)
	trajetoAsBytes, _ := json.Marshal(trajeto)
	veiculoTrajetoAsBytes, _ := json.Marshal(trajetoVeiculo)

	//Inserindo valor no Ledger
	idTrajeto := "trajeto-" + cdgUnicoTrajeto
	idUsuarioTrajeto := "veic_traj-" + cdgUnicoUsuarioTrajeto

	stub.PutState(idPlaca, veiculoAsBytesFinal)
	stub.PutState(idTrajeto, trajetoAsBytes)
	stub.PutState(idUsuarioTrajeto, veiculoTrajetoAsBytes)

	fmt.Println("Sucesso ao registrar trajeto")
	return shim.Success(nil)
}

func (s *SmartContract) registrarCarbono(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se arquivo py retornou 1 argumento
	if len(args) != 1 {
		return shim.Error("Era esperado 1 único argumento... Tente novamente!")
	}

	idPlaca := args[0]

	//Recuperando dados do usuário
	veiculoAsBytes, err := stub.GetState(idPlaca)
	if err != nil || veiculoAsBytes == nil {
		return shim.Error("Sua placa não existe.")
	}

	//Criando Struct para encapsular os dados do veiculo
	veiculo := Veiculo{}
	json.Unmarshal(veiculoAsBytes, &veiculo)

	idModelo := veiculo.IdCdgModelo

	//Recuperando dados do Modelo
	modeloAsBytes, err := stub.GetState(idModelo)
	if err != nil || modeloAsBytes == nil {
		return shim.Error("Esse modelo não existe.")
	}

	//Criando Struct para encapsular os dados dp modelo
	modelo := ModeloVeiculo{}
	json.Unmarshal(modeloAsBytes, &modelo)

	idFabricante := "fab-" + modelo.Marca

	//Recuperando dados do Faricante
	fabricanteAsBytes, err := stub.GetState(idFabricante)
	if err != nil || fabricanteAsBytes == nil {
		return shim.Error("Esse fabricante não existe.")
	}

	//Criando Struct para encapsular os dados dp modelo
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	disAcumuladorFloat, err := strconv.ParseFloat(veiculo.AcumuladorDistancia, 64)
	modEmissCo2Float, err := strconv.ParseFloat(modelo.EmissaoCo2, 64)

	co2Veiculo := disAcumuladorFloat * modEmissCo2Float
	co2VeiculoArredondado := Arredondar(co2Veiculo)

	fabCo2Float, err := strconv.ParseFloat(fabricante.Co2Tot, 64)

	fabCo2Float += co2VeiculoArredondado
	fabCO2String := fmt.Sprintf("%g", fabCo2Float)
	fabricante.Co2Tot = fabCO2String
	veiculo.AcumuladorDistancia = "0.0"

	fmt.Println("-----registro de carbono---")
	fmt.Println(veiculo)
	fmt.Println(modelo)
	fmt.Println(fabricante)

	//Encapsulando dados em arquivo JSON
	veiculoAsBytes, _ = json.Marshal(veiculo)
	modeloAsBytes, _ = json.Marshal(modelo)
	fabricanteAsBytes, _ = json.Marshal(fabricante)

	fmt.Println("-----BYTES-----")
	fmt.Println(veiculoAsBytes)
	fmt.Println(modeloAsBytes)
	fmt.Println(fabricanteAsBytes)

	stub.PutState(idPlaca, veiculoAsBytes)
	stub.PutState(idModelo, modeloAsBytes)
	stub.PutState(idFabricante, fabricanteAsBytes)

	fmt.Println("Sucesso ao computador co2 dos fabricantes")
	return shim.Success(nil)
}

func (s *SmartContract) registrarCredito(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se arquivo py retornou 1 argumento
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

	if fabricante.Co2Tot == "0.0" {
		fmt.Println("Não foi computado emissão de carbono para o fabricante: " + idFabricante)
		return shim.Success(nil)
	}

	fabCo2Float, err := strconv.ParseFloat(fabricante.Co2Tot, 64)

	var saldo = 50000.0 - fabCo2Float
	saldoString := fmt.Sprintf("%g", saldo)
	fabricante.SaldoCarbono = saldoString
	fabricante.Co2Tot = "0"

	fabricanteAsBytesFinal, _ := json.Marshal(fabricante)

	stub.PutState(idFabricante, fabricanteAsBytesFinal)

	fmt.Println("Saldo de carbono computado com sucesso: " + idFabricante)
	return shim.Success(nil)
}

func (s *SmartContract) anunciarOrdem(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se arquivo py retornou 3 argumento
	if len(args) != 3 {
		return shim.Error("Era esperado 3 único argumento... Tente novamente!")
	}

	nomeFabricante := args[0]
	tipoTransacao := args[1]
	saldoOferta := args[2]

	saldoOfertaFloat, err := strconv.ParseFloat(saldoOferta, 64)

	//Verificando se o fabricante realmente existe
	idFabricanteCompleto := "fab-" + nomeFabricante
	fabricanteAsBytes, err := stub.GetState(idFabricanteCompleto)
	if err != nil || fabricanteAsBytes == nil {
		return shim.Error("Seu fabricante não existe.")
	}

	//Criando Struct para encapsular os dados do fabricante
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	saldoCarbonoFloat, err := strconv.ParseFloat(fabricante.SaldoCarbono, 64)
	saldoFiduciarioFloat, err := strconv.ParseFloat(fabricante.SaldoFiduciario, 64)

	if tipoTransacao == "vender" {
		if saldoOfertaFloat > saldoCarbonoFloat {
			return shim.Error("Você não tem saldo de carbono suficiente")
		}

		saldoCarbonoFloat -= saldoOfertaFloat
		saldoCarbonoString := fmt.Sprintf("%g", saldoCarbonoFloat)
		fabricante.SaldoCarbono = saldoCarbonoString

	}

	if tipoTransacao == "comprar" {
		if saldoOfertaFloat > saldoFiduciarioFloat {
			return shim.Error("Você não tem saldo fiduciario suficiente")
		}
		saldoFiduciarioFloat -= saldoOfertaFloat
		saldoFiduciarioString := fmt.Sprintf("%g", saldoFiduciarioFloat)
		fabricante.SaldoFiduciario = saldoFiduciarioString
	}

	saldoOfertaString := fmt.Sprintf("%g", saldoOfertaFloat)

	ordemVenda := OrdemTransacao{
		IdComprador:       "0",
		ValorLance:        "0",
		StatusOrdem:       "recente",
		ProprietarioOrdem: idFabricanteCompleto,
		SaldoOfertado:     saldoOfertaString,
		TipoTransacao:     tipoTransacao,
	}

	ordemVendaAsBytes, _ := json.Marshal(ordemVenda)

	fmt.Println("-----------")
	fmt.Println(ordemVendaAsBytes)

	fabricanteAsBytesFinal, _ := json.Marshal(fabricante)

	fmt.Println("-----------")
	fmt.Println(fabricanteAsBytesFinal)

	idOrdem := "trans-" + nomeFabricante + Encode(AleatString(10))

	stub.PutState(idFabricanteCompleto, fabricanteAsBytesFinal)
	stub.PutState(idOrdem, ordemVendaAsBytes)

	fmt.Println("Ordem de " + tipoTransacao + " anunciado com sucesso!")
	return shim.Success(nil)
}

func (s *SmartContract) ordemLance(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	idTransacao := args[0]
	valorLance := args[1]
	idComprador := args[2]

	valorLanceFloat, err := strconv.ParseFloat(valorLance, 64)

	//Recuperando dados da transação
	ordemTransacaoAsBytes, err := stub.GetState(idTransacao)
	if err != nil || ordemTransacaoAsBytes == nil {
		return shim.Error("Seu proprietário não existe.")
	}

	//Recuperando dados do proprietário
	fabricanteAsBytes, err := stub.GetState(idComprador)
	if err != nil || ordemTransacaoAsBytes == nil {
		return shim.Error("Seu proprietário não existe.")
	}

	propietario := OrdemTransacao{}
	json.Unmarshal(ordemTransacaoAsBytes, &propietario)

	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	saldoFiat := fabricante.SaldoFiduciario
	saldoFiatFloat, err := strconv.ParseFloat(saldoFiat, 64)

	saldoCarb := fabricante.SaldoCarbono
	saldoCarbFloat, err := strconv.ParseFloat(saldoCarb, 64)

	if valorLanceFloat > saldoFiatFloat && propietario.TipoTransacao == "vender" {
		return shim.Error("Você não tem saldo fiduciario suficiente")
	}

	if valorLanceFloat > saldoCarbFloat && propietario.TipoTransacao == "comprar" {
		return shim.Error("Você não tem saldo de carbono suficiente")
	}

	propietario.StatusOrdem = "Andamento"
	propietario.ValorLance = valorLance
	propietario.IdComprador = idComprador

	ordemTransacaoAsBytesFinal, _ := json.Marshal(propietario)
	stub.PutState(idTransacao, ordemTransacaoAsBytesFinal)

	fmt.Println("Lance feito no sucesso")

	return shim.Success(nil)
}

func (s *SmartContract) fecharOrdem(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	idTransacao := args[0]

	//Recuperando dados da transação
	ordemTransacaoAsBytes, err := stub.GetState(idTransacao)
	if err != nil || ordemTransacaoAsBytes == nil {
		return shim.Error("Seu proprietário não existe.")
	}

	return shim.Success(nil)

}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
