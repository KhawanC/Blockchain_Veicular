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

type Veiculos struct {
	Veiculos []Veiculum `json:"Veiculo"`
}

type Veiculum struct {
	Categoria string `json:"Categoria"`
	Marca     string `json:"Marca"`
	Versao    string `json:"Versao"`
	Modelo    string `json:"Modelo"`
	Emissao   int    `json:"Emissao"`
	Codigo    string `json:"Codigo"`
	Placa     string `json:Placa`
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

	}

	return shim.Error("Chaincode não suporta essa função.")
}

func (s *SmartContract) registrarBanco(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	banco := args[0]

	var bancoInformacoes Veiculos

	bancoAsBytes, _ := json.Marshal(bancoInformacoes)

	stub.PutState(banco, bancoAsBytes)

	fmt.Println("Registrando seu banco de veiculos...")

	//notify procedure success
	return shim.Success(nil)
}

func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}
	banco := args[0]
	userPlaca := args[1]
	userEnum, err := strconv.Atoi(args[2])

	bancoAsBytes, err := stub.GetState(banco)

	if err != nil || bancoAsBytes == nil {
		return shim.Error("Erro na validação dos dados de veículos!")
	}

	bancoUsr := Veiculos{}
	json.Unmarshal(bancoAsBytes, &bancoUsr)

	bancoUsr.Veiculos[userEnum].Placa = userPlaca
	usrNormal := bancoUsr.Veiculos[userEnum]
	usrAsBytes, _ := json.Marshal(usrNormal)
	stub.PutState(usrNormal, usrAsBytes)

	fmt.Println("Registrando seu veículo...")

	return shim.Success(nil)

}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
