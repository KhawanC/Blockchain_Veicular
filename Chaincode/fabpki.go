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

type ModeloVeiculo struct {
	CdgModelo                               string  `json:"CdgModelo"` // PK
	Categoria                               string  `json:"Categoria"`
	Marca                                   string  `json:"Marca"`
	Versao                                  string  `json:"Versao"`
	Modelo                                  string  `json:"Modelo"`
	Motor                                   string  `json:"Motor"`
	Tipo_de_Propulsao                       string  `json:"Tipo_de_Propulsao"`
	Transmissao_Velocidades                 string  `json:"Transmissao_Velocidades_"`
	Ar_Condicionado                         string  `json:"Ar_Condicionado"`
	Direcao_Assistida                       string  `json:"Direcao_Assistida"`
	Combustivel                             string  `json:"Combustivel"`
	NMOG_Nox                                float64 `json:"NMOG_Nox"`
	CO                                      float64 `json:"CO"`
	CHO                                     float64 `json:"CHO"`
	Reducao_Relativa_Ao_Limite              string  `json:"Reducao_Relativa_Ao_Limite"`
	Etanol_CO2_Fossil                       float64 `json:"Etanol_CO2_Fossil"`
	Gasolina_Diesel_CO2_fossil              float64 `json:"Gasolina_Diesel_CO2_fossil"`
	Etanol_Cidade                           float64 `json:"Etanol_Cidade"`
	Etanol_Estrada                          float64 `json:"Etanol_Estrada"`
	Gasolina_Diesel_Eletrico_Cidade         float64 `json:"Gasolina_Diesel_Eletrico_Cidade"`
	Gasolina_Diesel_Eletrico_Estrada        float64 `json:"Gasolina_Diesel_Eletrico_Estrada"`
	Consumo_Energetico                      float64 `json:"Consumo_Energetico"`
	Classificacao_PBE_Relativo_na_Categoria string  `json:"Classificacao_PBE_Relativo_na_Categoria"`
	Classificação_PBE_Absoluto_Geral        string  `json:"Classificação_PBE_Absoluto_Geral"`
	Selo_CONPET_de_Eficiencia_Energetica    string  `json:"Selo_CONPET_de_Eficiencia_Energetica"`
}

type Usuario struct {
	Placa               string  `json:"Placa"`       // PK
	IdCdgModelo         string  `json:"IdCdgModelo"` //FK (Categoria)
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
	categoria := args[0]
	marca := args[1]
	versao := args[2]
	modelo := args[3]
	motor := args[4]
	tipo_de_Propulsao := args[5]
	transmissao_Velocidades := args[6]
	ar_Condicionado := args[7]
	direcao_Assistida := args[8]
	combustivel := args[9]
	nMOG_Nox := args[10]
	cO := args[11]
	cHO := args[12]
	reducao_Relativa_Ao_Limite := args[13]
	etanol_CO2_Fossil := args[14]
	gasolina_Diesel_CO2_fossil := args[15]
	etanol_Cidade := args[16]
	etanol_Estrada := args[17]
	gasolina_Diesel_Eletrico_Cidade := args[18]
	gasolina_Diesel_Eletrico_Estrada := args[19]
	consumo_Energetico := args[20]
	classificacao_PBE_Relativo_na_Categoria := args[21]
	classificação_PBE_Absoluto_Geral := args[22]
	celo_CONPET_de_Eficiencia_Energetica := args[23]

	NMOG_NoxFloat, err := strconv.ParseFloat(nMOG_Nox, 64)
	cOFloat, err := strconv.ParseFloat(cO, 64)
	cHOFloat, err := strconv.ParseFloat(cHO, 64)
	etanol_CO2_FossilFloat, err := strconv.ParseFloat(etanol_CO2_Fossil, 64)
	gasolina_Diesel_CO2_fossilFloat, err := strconv.ParseFloat(gasolina_Diesel_CO2_fossil, 64)
	etanol_CidadeFloat, err := strconv.ParseFloat(etanol_Cidade, 64)
	etanol_EstradaFloat, err := strconv.ParseFloat(etanol_Estrada, 64)
	gasolina_Diesel_Eletrico_CidadeFloat, err := strconv.ParseFloat(gasolina_Diesel_Eletrico_Cidade, 64)
	gasolina_Diesel_Eletrico_EstradaFloat, err := strconv.ParseFloat(gasolina_Diesel_Eletrico_Estrada, 64)
	consumo_EnergeticoFloat, err := strconv.ParseFloat(consumo_Energetico, 64)

	if err != nil {
		return shim.Error("Houve um problema ao converter o float")
	}

	var CategoriaInfor = ModeloVeiculo{ //Inserindo argumentos dentro da Struct Categoria
		CdgModelo:                               null,
		Categoria:                               categoria,
		Marca:                                   marca,
		Versao:                                  versao,
		Modelo:                                  modelo,
		Motor:                                   motor,
		Tipo_de_Propulsao:                       tipo_de_Propulsao,
		Transmissao_Velocidades:                 transmissao_Velocidades,
		Ar_Condicionado:                         ar_Condicionado,
		Direcao_Assistida:                       direcao_Assistida,
		Combustivel:                             combustivel,
		NMOG_Nox:                                NMOG_NoxFloat,
		CO:                                      cOFloat,
		CHO:                                     cHOFloat,
		Reducao_Relativa_Ao_Limite:              reducao_Relativa_Ao_Limite,
		Etanol_CO2_Fossil:                       etanol_CO2_FossilFloat,
		Gasolina_Diesel_CO2_fossil:              gasolina_Diesel_CO2_fossilFloat,
		Etanol_Cidade:                           etanol_CidadeFloat,
		Etanol_Estrada:                          etanol_EstradaFloat,
		Gasolina_Diesel_Eletrico_Cidade:         gasolina_Diesel_Eletrico_CidadeFloat,
		Gasolina_Diesel_Eletrico_Estrada:        gasolina_Diesel_Eletrico_EstradaFloat,
		Consumo_Energetico:                      consumo_EnergeticoFloat,
		Classificacao_PBE_Relativo_na_Categoria: classificacao_PBE_Relativo_na_Categoria,
		Classificação_PBE_Absoluto_Geral:        classificação_PBE_Absoluto_Geral,
		Selo_CONPET_de_Eficiencia_Energetica:    celo_CONPET_de_Eficiencia_Energetica,
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
