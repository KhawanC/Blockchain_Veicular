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
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Veiculo struct {
	Categoria string `json:"Categoria"`
	Marca     string `json:"Marca"`
	Versao    string `json:"Versao"`
	Modelo    string `json:"Modelo"`
	Emissao   int    `json:"Emissao"`
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

	if len(args) != 2 {
		return shim.Error("Não foi encontrado nenhum argumento. Tente novamente!")
	}

	idVeiculo := args[0]
	veiculoInfo := args[1]

	var info = Veiculo{veiculoInfo}

	fmt.Println("Registrando seu banco de veiculos...")

	//notify procedure success
	return shim.Success(nil)
}

func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	/*if len(args) != 2 {
		return shim.Error("Eram esperados 3 argumentos... Tente novamente!")
	}
	banco := args[0]
	userPlaca := args[1]
	userEnum, err := strconv.Atoi(args[2])

	bancoVirtual, _ := json.Marshal(banco)
	bancoAsBytes, err := stub.GetState(1)

	if err != nil || bancoAsBytes == nil || bancoVirtual != bancoAsBytes {
		return shim.Error("Erro na validação dos dados de veículos!")
	}

	MyBanco := Veiculos{}
	json.Unmarshal(bancoAsBytes, &MyBanco)

	MyBanco.Veiculos[userEnum].Placa = userPlaca
	userAsBytes, _ := json.Marshal(MyBanco.Veiculos[userEnum])

	stub.PutState(userPlaca, userAsBytes)
	fmt.Println("Registrando seu veículo...")

	return shim.Success(nil)*/

}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
