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

type SmartContract struct {
}

type ModeloVeiculo struct { //"model-"
	Categoria  string  `json:"Categoria"`
	Marca      string  `json:"Marca"`
	Versao     string  `json:"Versao"`
	Modelo     string  `json:"Modelo"`
	EmissaoCo2 float64 `json:"EmissaoCo2"`
}

type Veiculo struct { //"user-"
	IdCdgModelo         string  `json:"IdCdgModelo"` //FK (Categoria)
	AcumuladorDistancia float64 `json:"AcumuladorDistancia"`
}

type Trajeto struct { //"traj-"
	TrajetoDistancia float64 `json:"TrajetoDistancia"`
}

type TrajetoVeiculo struct { //"traj_user-"
	IdPlaca   string `json:"IdPlaca"`   //FK (Veiculo)
	IdTrajeto string `json:"IdTrajeto"` //FK (Trajeto)
}

type Fabricante struct { //"fab-""
	Co2Tot          float64 `json:"Co2_Tot"`
	SaldoCarbono    float64 `json:"SaldoCarbono"`
	SaldoFiduciario float64 `json:"Saldo_FIduciario"`
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

	emissFloat, err := strconv.ParseFloat(emissao_co2, 64)

	if err != nil {
		return shim.Error("Houve um problema ao converter o float")
	}

	var categoriaInfor = ModeloVeiculo{ //Inserindo argumentos dentro da Struct Categoria
		Categoria:  categoria,
		Marca:      marca,
		Versao:     versao,
		Modelo:     modelo,
		EmissaoCo2: emissFloat,
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

	var fabricanteInfor = Fabricante{
		Co2Tot:          0.0,
		SaldoCarbono:    0.0,
		SaldoFiduciario: 10000.0,
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
	if len(args) != 2 {
		return shim.Error("Eram esperados 2 argumentos... Tente novamente!")
	}

	userPlaca := args[0]
	CdgModeloUser := args[1]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Veiculo{
		IdCdgModelo:         CdgModeloUser,
		AcumuladorDistancia: 0.0,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)

	//Inserir valores no ledger. ID = placa do veículo
	idUserVeiculo := "veic-" + userPlaca
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

	distAcumulado := veiculo.AcumuladorDistancia //Inserindo o acumulador de distância do usuário dentro de uma variável

	distAcumulado += userDistFLoat //Adicionar a distancia do trajeto feito ao acumulador

	veiculo.AcumuladorDistancia = distAcumulado

	//Criar assinatura do trajeto
	trajeto.TrajetoDistancia = userDistFLoat

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

	idPlaca := args[0]

	//Recuperando dados do usuário
	userAsBytes, err := stub.GetState(idPlaca)
	if err != nil || userAsBytes == nil {
		return shim.Error("Sua placa não existe.")
	}

	//Criando Struct para encapsular os dados do veiculo
	veiculo := Veiculo{}
	json.Unmarshal(userAsBytes, &veiculo)

	idModelo := veiculo.IdCdgModelo

	//Recuperando dados do Modelo
	modeloAsBytes, err := stub.GetState(idModelo)
	if err != nil || modeloAsBytes == nil {
		return shim.Error("Esse modelo não existe.")
	}

	//Criando Struct para encapsular os dados dp modelo
	modelo := ModeloVeiculo{}
	json.Unmarshal(modeloAsBytes, &modelo)

	idFaricante := "fab-" + modelo.Marca

	//Recuperando dados do Faricante
	fabricanteAsBytes, err := stub.GetState(idFaricante)
	if err != nil || fabricanteAsBytes == nil {
		return shim.Error("Esse fabricante não existe.")
	}

	//Criando Struct para encapsular os dados dp modelo
	fabricante := Fabricante{}
	json.Unmarshal(fabricanteAsBytes, &fabricante)

	co2Veiculo := veiculo.AcumuladorDistancia * modelo.EmissaoCo2
	fabricante.Co2Tot += co2Veiculo
	veiculo.AcumuladorDistancia = 0.0

	//Encapsulando dados em arquivo JSON
	veiculoAsBytesFinal, _ := json.Marshal(veiculo)
	modeloAsBytesFinal, _ := json.Marshal(modelo)
	fabricanteAsBytesFinal, _ := json.Marshal(fabricante)

	stub.PutState(idPlaca, veiculoAsBytesFinal)
	stub.PutState(idModelo, modeloAsBytesFinal)
	stub.PutState(idFaricante, fabricanteAsBytesFinal)

	fmt.Println("Sucesso ao computador co2 dos fabricantes")
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
