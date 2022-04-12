package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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
	Placa     string `json:"Placa"`
}

func main() {

	bancoAsBytes, err := ioutil.ReadFile("dadosVeiculares.json")
	if err != nil {
		log.Fatal(err)
	}
	userPlaca := "ABC1D23"
	userEnum := 4

	/*var veiculo Veiculos
	json.Unmarshal([]byte(banco), &veiculo)

	dadosUsuario := [2]string{"4", "ABC12D3"}
	usuarioEnum, err := strconv.Atoi(dadosUsuario[0])
	usuarioPlaca := dadosUsuario[1]

	veiculo.Veiculos[usuarioEnum].Placa = usuarioPlaca
	usrFinal := veiculo.Veiculos[usuarioEnum]

	veiculoo, _ := json.Marshal(usrFinal)

	fmt.Println(string(veiculoo))*/

	MyBanco := Veiculos{}
	json.Unmarshal(bancoAsBytes, &MyBanco)

	MyBanco.Veiculos[userEnum].Placa = userPlaca

	fmt.Println(MyBanco.Veiculos[userEnum])

}
