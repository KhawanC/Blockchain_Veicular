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

type Categoria struct {
	CdgCategoria string `json:"CdgCategoria"` // PK
	Categoria    string `json:"Categoria"`
	Marca        string `json:"Marca"`
	Versao       string `json:"Versao"`
	Modelo       string `json:"Modelo"`
	EmissaoPad   string `json:"EmissaoPad"`
}

type Usuario struct {
	Placa               string `json:"Placa"`          // PK
	IdCdgCategoria      string `json:"IdCdgCategoria"` //FK (Categoria)
	AcumuladorDistancia string `json:"AcumuladorDistancia"`
	Co2Emitido          string `json:"Co2Emitido"`
	CreditosDeCarbono   string `json:"CreditosDeCarbono"`
}

type Trajeto struct {
	TrajetoHash      string `json:"TrajetoHash"` // PK
	TrajetoDistancia string `json:"TrajetoDistancia"`
}

type TrajetoUsuario struct {
	Viagem    string `json:"Viagem"`    // PK
	IdPlaca   string `json:"IdPlaca"`   //FK (Usuario)
	IdTrajeto string `json:"IdTrajeto"` //FK (Trajeto)
}

type Token struct {
	CreditoToken  string `json:"CreditoToken"` //PK
	MetaDeEmissao string `json:"MetaDeEmissao"`
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

	} else if fn == "calcularCreditos" {
		return s.calcularCreditos(stub, args)
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

	//Arredondar a emissao padrao para 2 casas decimais
	emissFloat, err := strconv.ParseFloat(emissao, 64)
	emissString := fmt.Sprintf("%.2f", emissFloat)

	//Criando o token
	var credito = Token{
		CreditoToken: "Credito_Token",
	}

	//Inserindo argumentos dentro da Struct Categoria
	var CategoriaInfor = Categoria{
		Categoria:    categoria,
		Marca:        marca,
		Versao:       versao,
		Modelo:       modelo,
		EmissaoPad:   emissString,
		CdgCategoria: codigo,
	}

	//Encapsulando as informações em arquivo JSON
	CategoriaAsBytes, _ := json.Marshal(CategoriaInfor)
	tokenAsBytes, _ := json.Marshal(credito)

	//Inserindo valores no ledger, com uma informação associada à uma chave
	stub.PutState(codigo, CategoriaAsBytes)
	stub.PutState(credito.CreditoToken, tokenAsBytes)

	if err != nil {
		return shim.Error("Houve um problema ao converter o float")
	}

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
	cdgCategoriaUser := args[1]

	//Criar Struct para manipular as informações do veículo
	userVeiculo := Usuario{
		Placa:               userPlaca,
		IdCdgCategoria:      cdgCategoriaUser,
		AcumuladorDistancia: "0.0",
		Co2Emitido:          "0.0",
		CreditosDeCarbono:   "0.0",
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

	userDistFLoat, err := strconv.ParseFloat(userDistancia, 64)

	//Criando código unico para Struct trajeto
	cdgUnico := Encode(AleatString(20))

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

	//Converter a distância acumulada do usuario para Float
	distAcumulFLoat, err := strconv.ParseFloat(usuario.AcumuladorDistancia, 64)

	//Adicionar a distancia do trajeto feito ao acumulador
	distAcumulFLoat += userDistFLoat

	//Converter valor em FLoat do acumulador para String e inserir no objeto do usuário
	distAcumulString := fmt.Sprintf("%.2f", distAcumulFLoat)
	usuario.AcumuladorDistancia = distAcumulString

	//Criar assinatura do trajeto
	trajeto.TrajetoHash = cdgUnico
	trajeto.TrajetoDistancia = userDistancia

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

func (s *SmartContract) calcularMeta(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("O argumeto esperado não foi encontrado")
	}

	mediaEmissPad := args[0]
	mediaTrajetos := args[1]

	tokenAsBytes, err := stub.GetState("Credito_Token")
	if err != nil {
		return shim.Error("Nao foi possível recuperar o token da rede")
	}

	token := Token{}

	//Convertendo os valores de String para Float
	mediaEmissFloat, err2 := strconv.ParseFloat(mediaEmissPad, 64)
	mediaTrajetoFloat, err2 := strconv.ParseFloat(mediaTrajetos, 64)

	if err2 != nil {
		return shim.Error("Nao foi possível converter algum dos valores de String para Float")
	}

	metaDeEmissFloat := (mediaEmissFloat * mediaTrajetoFloat)
	metaDeEmissaoString := fmt.Sprintf("%.2f", metaDeEmissFloat)
	token.MetaDeEmissao = metaDeEmissaoString

	tokenAsBytes, _ = json.Marshal(token)
	stub.PutState("Credito_Token", tokenAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) calcularCreditos(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("O argumeto esperado não foi encontrado")
	}

	fmt.Println("Iniciando calculo de créditos")
	placa := args[0]

	tokenAsBytes, err := stub.GetState("Credito_Token")
	if err != nil {
		return shim.Error("Nao foi possível recuperar o token da rede")
	}

	token := Token{}
	json.Unmarshal(tokenAsBytes, &token)

	userAsBytes, err1 := stub.GetState(placa)
	if err1 != nil || userAsBytes == nil {
		return shim.Error("Nao foi possível recuperar o seu veículo no sistema")
	}

	usuario := Usuario{}
	json.Unmarshal(userAsBytes, &usuario)

	fmt.Println("Sua placa é: " + usuario.Placa)

	categoriaAsBytes, err2 := stub.GetState(usuario.IdCdgCategoria)
	if err2 != nil || categoriaAsBytes == nil {
		return shim.Error("Não foi possível recuperar as informações da categoria")
	}

	categoria := Categoria{}
	json.Unmarshal(categoriaAsBytes, &categoria)

	fmt.Println("Esse veículo é um " + categoria.Marca)

	//Converter acumulador de String para Float
	distAcumulFLoat, err := strconv.ParseFloat(usuario.AcumuladorDistancia, 64)

	if distAcumulFLoat == 0 {
		return shim.Error("O veículo não rodou durante o mês")
	}

	usuario.Co2Emitido = "0.0"

	//Converter Meta de Emissão do Token
	metaEmissToken, err := strconv.ParseFloat(token.MetaDeEmissao, 64)

	//Converter emissao padrão da categoria de String para Float
	emissPadFLoat, err := strconv.ParseFloat(categoria.EmissaoPad, 64)

	//Realizar calculo de Emissão
	co2EmitidoFloat := emissPadFLoat * distAcumulFLoat

	//Converter Emissão para String e inserir no valor respectivo do veículo
	co2EmitidoString := fmt.Sprintf("%.2f", co2EmitidoFloat)
	usuario.Co2Emitido = co2EmitidoString

	//Converter creditos já existentes para FLoat
	creditosOld, err := strconv.ParseFloat(usuario.CreditosDeCarbono, 64)

	//Calcular créditos, zerar valor de emissão
	creditos := metaEmissToken - co2EmitidoFloat
	creditosNew := creditosOld + creditos

	creditosString := fmt.Sprintf("%.2f", creditosNew)

	usuario.CreditosDeCarbono = creditosString
	usuario.AcumuladorDistancia = "0"

	if err != nil {
		return shim.Error("Houve um problema ao converter o float")
	}

	userAsBytes, _ = json.Marshal(usuario)
	stub.PutState(placa, userAsBytes)

	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Erro ao compilar Smart Contract: %s\n", err)
	}
}
