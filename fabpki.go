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

//Criar sequencia de letras
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func AleatString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Encode(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

type SmartContract struct {
}

type Veiculo struct {
	CdgVeiculo string `json:"CdgVeiculo"` // PK
	Categoria  string `json:"Categoria"`
	Marca      string `json:"Marca"`
	Versao     string `json:"Versao"`
	Modelo     string `json:"Modelo"`
	EmissaoPad string `json:"EmissaoPad"`
}

type Usuario struct {
	Placa               string  `json:"Placa"`        // PK
	IdCdgVeiculo        string  `json:"IdCdgVeiculo"` //FK (Veiculo)
	AcumuladorDistancia float64 `json:"AcumuladorDistancia"`
	CreditosDeCarbono   float64 `json:"CreditosDeCarbono"`
	QtdViagens          float64 `json:"qtdViagens"`
}

type Trajeto struct {
	TrajetoHash      string  `json:"TrajetoHash"` // PK
	TrajetoDistancia float64 `json:"TrajetoDistancia"`
}

type TrajetoUsuario struct {
	Viagem    string `json:"Viagem"`      // PK
	IdPlaca   string `json:"IdPlaca"`     //FK (Usuario)
	IdTrajeto string `json:"TrajetoHash"` //FK (Trajeto)
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
		Placa:               userPlaca,
		IdCdgVeiculo:        cdgVeiculoUser,
		AcumuladorDistancia: 0.0,
		QtdViagens:          0.0,
	}

	veiculoAsBytes, _ := json.Marshal(userVeiculo)

	//Inserir valores no ledger. ID = placa do veículo
	stub.PutState(userPlaca, veiculoAsBytes)

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

	//Criando código unico para Struct trajeto
	cdgUnico := Encode(AleatString(20))

	//Converter distância do argumento para Float64 pois ela veio como String
	distFloat, err := strconv.ParseFloat(userDistancia, 64)
	if err != nil {
		return shim.Error("Erro ao converter distância do usuário")
	}

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

	//Atualizando dados de trajeto do usuario
	trajetoOld := usuario.AcumuladorDistancia
	usuario.AcumuladorDistancia = trajetoOld + distFloat
	viagensOld := usuario.QtdViagens
	usuario.QtdViagens = viagensOld + 1.0

	//Criar assinatura do trajeto
	trajeto.TrajetoHash = cdgUnico
	trajeto.TrajetoDistancia = distFloat

	//Associar trajeto com o usuário
	trajetoUsuario.Viagem = trajeto.TrajetoHash + "-" + usuario.Placa
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
