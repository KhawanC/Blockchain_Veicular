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
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Veiculo struct {
	Categoria  string `json:"Categoria"`
	Marca      string `json:"Marca"`
	Versao     string `json:"Versao"`
	Modelo     string `json:"Modelo"`
	EmissaoPad string `json:"EmissaoPad"`
	CdgVeiculo string `json:"CdgVeiculo"`
}

type Usuario struct {
	Placa        string `json:"Placa"`
	IdCdgVeiculo string `json:"IdCdgVeiculo"`
}

type TrajetoUsuario struct {
	IdPlaca          string  `json:"IdPlaca"`
	TrajetoAcumulado float64 `json:"TrajetoAcumulador"`
	QtdTrajetos      int     `json:"QtdTrajetos"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registrarBanco" {
		return s.registrarBanco(stub, args)

	} else if fn == "registrarUsuario" {
		return s.registrarUsuario(stub, args)

	} else if fn == "registrarTrajeto" {
		return s.registrarTrajeto(stub, args)
	}

	return shim.Error("Chaincode não suporta essa função.")
}

//Função que recebe bancoLedger.py e inserer o banco de dados com os veículos
func (s *SmartContract) registrarBanco(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificando se a quantidade de argumnetos é maior que 6
	if len(args) != 6 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	//Inserindo argumentos dentro de variáveis
	codigo := args[0]
	categoria := args[1]
	marca := args[2]
	versao := args[3]
	modelo := args[4]
	emissao := args[5]

	//Inserindo argumentos dentro da Struct Veiculo
	var veiculoInfor = Veiculo{
		Categoria:  categoria,
		Marca:      marca,
		Versao:     versao,
		Modelo:     modelo,
		EmissaoPad: emissao,
		CdgVeiculo: codigo,
	}

	//Encapsulando as informações do veículo em formato JSON
	veiculoAsBytes, _ := json.Marshal(veiculoInfor)

	//Inserindo valores no ledger, com uma informação associada à uma chave
	stub.PutState(codigo, veiculoAsBytes)

	//Confirmação do chaincode
	fmt.Println("Registrando seu banco de veiculos...")
	return shim.Success(nil)
}

//Função que recebe userLedger e insere o veículo do usuário no Ledger
func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se existem mais de 2 argumentos no código do cliente
	if len(args) != 2 {
		return shim.Error("Eram esperados 2 argumentos... Tente novamente!")
	}
	userPlaca := args[0]
	cdgVeiculoUser := args[1]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Usuario{
		Placa:        userPlaca,
		IdCdgVeiculo: cdgVeiculoUser,
	}

	veiculoAsBytesFinal, _ := json.Marshal(userVeiculo)

	//Inserir valores no ledger. ID = placa do veículo
	stub.PutState(userPlaca, veiculoAsBytesFinal)

	return shim.Success(nil)

}

func (s *SmartContract) registrarTrajeto(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se arquivo py retornou 2 argumentos
	if len(args) != 2 {
		return shim.Error("Eram esperados 2 argumentos... Tente novamente!")
	}

	//Dar nome aos argumentos
	userPlaca := args[0]
	userDistancia := args[1]

	//Converter distância do argumento para Float64 pois ela veio como String
	distFloat, err := strconv.ParseFloat(userDistancia, 64)
	if err != nil {
		return shim.Error("Erro ao converter distância do usuário")
	}

	//Criar Struct do trajeto e do usuário
	infoTrajeto := TrajetoUsuario{}
	userVeiculo := Usuario{}

	//Verificar se o usuário existe no Ledger
	UsuarioAsBytes, err := stub.GetState(userPlaca)
	if err != nil || UsuarioAsBytes == nil {

		//Caso não exista, criar uma assinatura no Ledger com a sua placa, mas suas especificações ficarão desconhecidas
		userVeiculo.Placa = userPlaca

		//Inserindo informações do trajendo em um acumulador e incrementando 1 ao seu contador de trajetos
		infoTrajeto.IdPlaca = userPlaca
		infoTrajeto.TrajetoAcumulado += distFloat
		infoTrajeto.QtdTrajetos += 1

		//Encapsulando informações em formato JSON
		TrajetoAsBytesFinal, _ := json.Marshal(infoTrajeto)
		UserAsBytesFinal, _ := json.Marshal(userVeiculo)

		//Inserindo informações no Ledger
		stub.PutState(userPlaca+strconv.Itoa(infoTrajeto.QtdTrajetos), TrajetoAsBytesFinal)
		stub.PutState(userPlaca, UserAsBytesFinal)

		return shim.Success(nil)
	}

	//Caso exista, apenas criar uma assinatura de trajeto
	infoTrajeto.IdPlaca = userPlaca
	infoTrajeto.TrajetoAcumulado += distFloat
	infoTrajeto.QtdTrajetos += 1

	//Encapsulando informações em formato JSON
	trajetoAsBytesFinal, _ := json.Marshal(infoTrajeto)

	//Inserindo informações no Ledger
	stub.PutState(userPlaca, trajetoAsBytesFinal)

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
