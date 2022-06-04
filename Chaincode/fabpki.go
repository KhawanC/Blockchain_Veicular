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

type Categoria struct {
	CdgCategoria string  `json:"CdgCategoria"` // PK
	Categoria    string  `json:"Categoria"`
	Marca        string  `json:"Marca"`
	Versao       string  `json:"Versao"`
	Modelo       string  `json:"Modelo"`
	EmissaoPad   float64 `json:"EmissaoPad"`
}

type Usuario struct {
	Placa               string  `json:"Placa"`          // PK
	IdCdgCategoria      string  `json:"IdCdgCategoria"` //FK (Categoria)
	AcumuladorDistancia float64 `json:"AcumuladorDistancia"`
	Co2Emitido          float64 `json:"Co2Emitido"`
}

type Trajeto struct {
	TrajetoHash      string  `json:"TrajetoHash"` // PK
	TrajetoDistancia float64 `json:"TrajetoDistancia"`
}

type TrajetoUsuario struct {
	Viagem    string `json:"Viagem"`    // PK
	IdPlaca   string `json:"IdPlaca"`   //FK (Usuario)
	IdTrajeto string `json:"IdTrajeto"` //FK (Trajeto)
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

//Função que recebe ./bancoLedger.py e inserer o banco de dados com os veículos
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

	emissFloat, err := strconv.ParseFloat(emissao, 64) //Convertendo a emissão para float

	if err != nil {
		return shim.Error("Houve um problema ao converter o float")
	}

	var CategoriaInfor = Categoria{ //Inserindo argumentos dentro da Struct Categoria
		Categoria:    categoria,
		Marca:        marca,
		Versao:       versao,
		Modelo:       modelo,
		EmissaoPad:   emissFloat,
		CdgCategoria: codigo,
	}

	CategoriaAsBytes, _ := json.Marshal(CategoriaInfor) //Encapsulando as informações em arquivo JSON

	stub.PutState(codigo, CategoriaAsBytes) //Inserindo valores no ledger, com uma informação associada à uma chave

	//Confirmação do chaincode
	fmt.Println("Registro de categoria inserido com sucesso")
	return shim.Success(nil)
}

//Função que recebe "./userLedger" e insere o veículo do usuário no Ledger
func (s *SmartContract) registrarUsuario(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Verificar se existem mais de 2 argumentos no código do cliente
	if len(args) != 2 {
		return shim.Error("Eram esperados 2 argumentos... Tente novamente!")
	}

	userPlaca := args[0]
	cdgCategoriaUser := args[1]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Usuario{
		Placa:               userPlaca,
		IdCdgCategoria:      cdgCategoriaUser,
		AcumuladorDistancia: 0.0,
		Co2Emitido:          0.0,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)

	//Inserir valores no ledger. ID = placa do veículo
	stub.PutState(userPlaca, veiculoAsBytes)

	//Confirmação do chaincode
	fmt.Println("Registro de usuário inserido com sucesso")
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

	userDistFLoat, err := strconv.ParseFloat(userDistancia, 64) //Convertendo a distância recebida para float

	cdgUnico := Encode(AleatString(20)) //Criando código unico para Struct trajeto

	//Criar Struct do trajeto e do usuário
	trajetoUsuario := TrajetoUsuario{}
	trajeto := Trajeto{}

	//Verificar se o usuário existe no Ledger
	UsuarioAsBytes, err := stub.GetState(userPlaca)
	if err != nil || UsuarioAsBytes == nil {
		return shim.Error("Sua placa não existe")
	}

	fmt.Println("Informações do usuário obtidas")

	//Convertendo as informações do usuário em um objeto
	usuario := Usuario{}
	json.Unmarshal(UsuarioAsBytes, &usuario)

	distAcumulado := usuario.AcumuladorDistancia //Inserindo o acumulador de distância do usuário dentro de uma variável

	distAcumulado += userDistFLoat //Adicionar a distancia do trajeto feito ao acumulador

	usuario.AcumuladorDistancia = distAcumulado

	//Criar assinatura do trajeto
	trajeto.TrajetoHash = cdgUnico
	trajeto.TrajetoDistancia = userDistFLoat

	//Associar trajeto com o usuário
	trajetoUsuario.Viagem = "travel-" + trajeto.TrajetoHash + "-" + usuario.Placa
	trajetoUsuario.IdPlaca = usuario.Placa
	trajetoUsuario.IdTrajeto = trajeto.TrajetoHash

	//Encapsulando dados em arquivo JSON
	UsuarioAsBytesFinal, _ := json.Marshal(usuario)
	TrajetoAsBytes, _ := json.Marshal(trajeto)
	UsuarioTrajetoAsBytes, _ := json.Marshal(trajetoUsuario)

	//Inserindo valor no Ledger
	stub.PutState(usuario.Placa, UsuarioAsBytesFinal)
	stub.PutState(trajeto.TrajetoHash, TrajetoAsBytes)
	stub.PutState(trajetoUsuario.Viagem, UsuarioTrajetoAsBytes)

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
