package main

import (
	"fmt"
	"strconv"
)

type Veiculo struct {
	Categoria string  `json:"Categoria"`
	Marca     string  `json:"Marca"`
	Versao    string  `json:"Versao"`
	Modelo    string  `json:"Modelo"`
	Emissao   float64 `json:"Emissao"`
	Placa     string  `json:"Placa"`
}

func main() {
	veiculoInfo := Veiculo{Categoria: "Compacto", Marca: "VW", Versao: "GOL", Modelo: "PATRULHEIRO", Emissao: 113.08668705501557}

	numeroString := "113.08668705501557"

	if numeroDouble, err := strconv.ParseFloat(numeroString, 32); err == nil {
		fmt.Println(numeroDouble)
	}

	if numeroDouble, err := strconv.ParseFloat(numeroString, 64); err == nil {
		fmt.Println(numeroDouble)
	}

	fmt.Println(veiculoInfo)

}
