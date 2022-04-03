package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	//"bytes"
	//"crypto/rand"
	//"crypto/rsa"
	//"crypto/x509"
	//"encoding/base64"
	//"encoding/pem"
	//"log"
	//"math/big"
	//"unicode/utf8"
)

type Veiculo_Cif struct {
	Placa     string `json:"Placa"`
	Categoria string `json:"Categoria"`
	Codigo    string `json:"Codigo"`
}

type Veiculo struct {
	Categoria        string    `json:"Categoria"`
	Marca            string    `json:"Marca"`
	Versao           string    `json:"Versao"`
	Modelo           string    `json:"Modelo"`
	Emissao          int       `json:"Emissao"`
	Codigo           string    `json:"Codigo"`
	Veiculo_Compacto []Veiculo `json:"Veiculo_Compacto"`
	Veiculo_Medio    []Veiculo `json:"Veiculo_Medio"`
}

func verificarErro(err error) {
	if err != nil {
		panic(err)
	}
}

func abrirJson(arquivo string) *os.File {
	f, err := os.Open(arquivo)
	verificarErro(err)
	return f
}

func main() {

	tokenFile := abrirJson("TOKEN.json")
	tabelaFile := abrirJson("dadosVeiculares.json")
	defer tokenFile.Close()
	defer tabelaFile.Close()

	tokenAsBytes, _ := ioutil.ReadAll(tokenFile)
	tabelaAsBytes, _ := ioutil.ReadAll(tabelaFile)

	var informacoes_tabela Veiculo
	var informacoes_token Veiculo_Cif

	json.Unmarshal(tabelaAsBytes, &informacoes_tabela)
	json.Unmarshal(tokenAsBytes, &informacoes_token)

	arg1 := os.Args[1]
	if arg1 == informacoes_token.Placa {
	} else {
		panic("SUA PLACA NÃO COINCIDE")
	}

	if informacoes_token.Categoria == "Veiculo_Medio" {
		for i := 0; i <= len(informacoes_tabela.Veiculo_Medio)-1; i++ {
			if informacoes_token.Codigo == informacoes_tabela.Veiculo_Medio[i].Codigo {
				ledger_map := make(map[string]string)
				ledger_map["Placa"] = informacoes_token.Placa
				ledger_map["Versao"] = informacoes_tabela.Veiculo_Medio[i].Versao
				ledger_map["Modelo"] = informacoes_tabela.Veiculo_Medio[i].Modelo
				ledger_map["Emissao"] = strconv.Itoa(informacoes_tabela.Veiculo_Medio[i].Emissao)
				ledger_map["Codigo"] = informacoes_tabela.Veiculo_Medio[i].Codigo

				ledger_map2, _ := json.Marshal(ledger_map)
				fmt.Println(string(ledger_map2))
			}
		}
	}
	if informacoes_token.Categoria == "Veiculo_Compacto" {
		for i := 0; i <= len(informacoes_tabela.Veiculo_Compacto)-1; i++ {
			if informacoes_token.Codigo == informacoes_tabela.Veiculo_Compacto[i].Codigo {
				ledger_map := make(map[string]string)
				ledger_map["Placa"] = informacoes_token.Placa
				ledger_map["Versao"] = informacoes_token.Placa
				ledger_map["Modelo"] = informacoes_tabela.Veiculo_Compacto[i].Modelo
				ledger_map["Emissao"] = strconv.Itoa(informacoes_tabela.Veiculo_Compacto[i].Emissao)
				ledger_map["Codigo"] = informacoes_tabela.Veiculo_Compacto[i].Codigo

				ledger_map2, _ := json.Marshal(ledger_map)
				fmt.Println(string(ledger_map2))
			}
		}
	}

	//CODIGO ANTIGO, TALVEZ EU USE
	/*mess64 := base64.StdEncoding.EncodeToString([]byte(message))
	privPem, _ := readFile("keys/privkey.pem")
	//pubPem, _ := readFile("keys/pubkey.pem")
	der, _ := pem.Decode(privPem)

	if der == nil {
		log.Fatal("Não há chave privada")
	}

	privKey, err := x509.ParsePKCS1PrivateKey(der.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	textoReal, err := base64.StdEncoding.DecodeString(mess64)
	if err != nil {
		log.Fatal(err)
	}
	c := new(big.Int).SetBytes(textoReal)
	textoDecifrado := c.Exp(c, (*big.Int)(privKey.D), (*big.Int)(privKey.N)).Bytes()

	newTexto := bytes.NewBuffer(textoDecifrado)

	fmt.Println(utf8.DecodeLastRuneInString(string(textoDecifrado)))

	textoDecifrado, err := rsa.DecryptPKCS1v15(rand.Reader, privKeys, message)

	fmt.Printf("TEXTO: %x %T\n", messBase64, messBase64)
	rsaCiphertext, _ := hex.DecodeString("aabbccddeeff")
	messData := bytes.NewBuffer(message)
	fmt.Printf("priv tipo: %T", privKeyData)
	fmt.Printf("mess tipo: %T", messData)*/

}