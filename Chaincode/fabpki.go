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
	CdgModelo   string  `json:"CdgModelo"` // PK
	Categoria   string  `json:"Categoria"`
	Marca       string  `json:"Marca"`
	Versao      string  `json:"Versao"`
	Modelo      string  `json:"Modelo"`
	Emissao_Co2 float64 `json:"Gasolina_Diesel_Eletrico_Estrada"`
}

type Usuario struct { //"user-"
	Placa               string  `json:"Placa"`       // PK
	IdCdgModelo         string  `json:"IdCdgModelo"` //FK (Categoria)
	AcumuladorDistancia float64 `json:"AcumuladorDistancia"`
	Co2Emitido          float64 `json:"Co2Emitido"`
}

type Trajeto struct { //"traj-"
	TrajetoHash      string  `json:"TrajetoHash"` // PK
	TrajetoDistancia float64 `json:"TrajetoDistancia"`
}

type TrajetoUsuario struct { //"traj_user-"
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

	var CategoriaInfor = ModeloVeiculo{ //Inserindo argumentos dentro da Struct Categoria
		CdgModelo:   cdg,
		Categoria:   categoria,
		Marca:       marca,
		Versao:      versao,
		Modelo:      modelo,
		Emissao_Co2: emissFloat,
	}

	CategoriaAsBytes, _ := json.Marshal(CategoriaInfor) //Encapsulando as informações em arquivo JSON

	idCdgLedger := "model-" + cdg

	stub.PutState(idCdgLedger, CategoriaAsBytes) //Inserindo valores no ledger, com uma informação associada à uma chave

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
	CdgModeloUser := args[1]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Usuario{
		Placa:               userPlaca,
		IdCdgModelo:         CdgModeloUser,
		AcumuladorDistancia: 0.0,
		Co2Emitido:          0.0,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)

	//Inserir valores no ledger. ID = placa do veículo
	idUserVeiculo := "user-" + userPlaca
	stub.PutState(idUserVeiculo, veiculoAsBytes)

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

	cdgUnico := Encode(AleatString(20))                         //Criando código unico para Struct trajeto
	userDistFLoat, err := strconv.ParseFloat(userDistancia, 64) //Convertendo a distância recebida para float

	//Recuperando dados do usuário
	userAsBytes, err := stub.GetState(userPlaca)
	if err != nil || userAsBytes == nil {
		return shim.Error("Sua placa não existe.")
	}

	//Criando Struct para encapsular os dados
	usuario := Usuario{}
	json.Unmarshal(userAsBytes, &usuario)

	//Criar Struct do trajeto e do usuário
	trajetoUsuario := TrajetoUsuario{}
	trajeto := Trajeto{}

	distAcumulado := usuario.AcumuladorDistancia //Inserindo o acumulador de distância do usuário dentro de uma variável

	distAcumulado += userDistFLoat //Adicionar a distancia do trajeto feito ao acumulador

	usuario.AcumuladorDistancia = distAcumulado

	//Criar assinatura do trajeto
	trajeto.TrajetoHash = "traj-" + cdgUnico
	trajeto.TrajetoDistancia = userDistFLoat

	//Associar trajeto com o usuário
	trajetoUsuario.Viagem = "traj_user-" + trajeto.TrajetoHash + "-" + usuario.Placa
	trajetoUsuario.IdPlaca = usuario.Placa
	trajetoUsuario.IdTrajeto = trajeto.TrajetoHash

	//Encapsulando dados em arquivo JSON
	UsuarioAsBytesFinal, _ := json.Marshal(usuario)
	TrajetoAsBytes, _ := json.Marshal(trajeto)
	UsuarioTrajetoAsBytes, _ := json.Marshal(trajetoUsuario)

	//Inserindo valor no Ledger
	stub.PutState(userPlaca, UsuarioAsBytesFinal)
	stub.PutState(trajeto.TrajetoHash, TrajetoAsBytes)
	stub.PutState(trajetoUsuario.Viagem, UsuarioTrajetoAsBytes)

	fmt.Println("Sucesso ao registrar trajeto")

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}