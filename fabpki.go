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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var idBanco = "1"

type SmartContract struct {
}

type Veiculo struct {
	Categoria  string `json:"Categoria"`
	Marca      string `json:"Marca"`
	Versao     string `json:"Versao"`
	Modelo     string `json:"Modelo"`
	Emissao    string `json:"Emissao"`
	Codigo     string `json:"Codigo"`
	Placa      string `json:"Placa"`
	EmissAcum  string `json:"EmissAcum"`
	Registrado bool   `json:"Registrado"`
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
	idVeiculo := args[0]
	categoria := args[1]
	marca := args[2]
	versao := args[3]
	modelo := args[4]
	emissao := args[5]
	codigo := args[0]

	//Inserindo argumentos dentro da Struct Veiculo
	var veiculoInfor = Veiculo{
		Categoria:  categoria,
		Marca:      marca,
		Versao:     versao,
		Modelo:     modelo,
		Emissao:    emissao,
		Codigo:     codigo,
		Registrado: true,
	}

	//Encapsulando as informações do veículo em formato JSON
	veiculoAsBytes, _ := json.Marshal(veiculoInfor)

	//Inserindo valores no ledger, com uma informação associada à uma chave
	stub.PutState(idVeiculo, veiculoAsBytes)

	//Confirmação do chaincode
	fmt.Println("Registrando seu banco de veiculos...")
	return shim.Success(nil)
}

//Função que recebe userLedger e insere o veículo do usuário no Ledger
func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se existem mais de 2 argumentos no código do cliente
	if len(args) != 2 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}
	userPlaca := args[0]
	cdgVeiculoUser := args[1]

	//Buscar informações referentes ao código do veículo do usuário dentro do ledger
	veiculoAsBytes, err := stub.GetState(cdgVeiculoUser)
	if err != nil || veiculoAsBytes == nil {
		return shim.Error("Erro ao receber dados do veículo")
	}

	//Criar um Struct para manipular as informações do veículo
	userVeiculo := Veiculo{}

	//Convertendo veiculoAsBytes para Struct do veículo
	json.Unmarshal(veiculoAsBytes, &userVeiculo)

	//Atualizar placa vazia da Struct com a placa do usuário
	//Como as informações do veículo já vieram com o Struct userVeiculo, vou alterar apenas a placa
	userVeiculo.Placa = userPlaca

	//Inserir valores no ledger. ID = placa do veículo
	stub.PutState(userPlaca, userVeiculo)

	return shim.Success(nil)

}

func (s *SmartContract) registrarTrajeto(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
