package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Veiculo struct {
	Placa       string `json:"placa"`
	Combustivel string `json:"combustivel"`
	Categoria   string `json:"categoria"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func registrarCliente(v1, v2, v3 string) map[string]string {

	numPlaca := v1
	tipoCategoria := v2
	tipoCombustivel := v3

	registro := make(map[string]string)

	registro["placa"] = numPlaca
	registro["categoria"] = tipoCategoria
	registro["combustivel"] = tipoCombustivel

	var veiculo = Veiculo{Placa: registro["placa"], Categoria: registro["categoria"], Combustivel: registro["combustivel"]}
	var veiculoAsBytes, _ = json.Marshal(veiculo)

	err := os.WriteFile("veiculo.json", veiculoAsBytes, 0644)
	check(err)

	fmt.Printf("Arquivo criada com seus parâmetros:\n\n%v\n\n%v", veiculo, veiculoAsBytes)

	return registro
}

func main() {
	registrarCliente("KVK1234", "2", "4")

	/*var leitor = "KVK23241"
	var registro [3]string
	registro[0] = "KVK23241"
	registro[1] = "2"
	registro[2] = "4"
	registrarCliente(registro)
	calcularCarbono(leitor)

	var veiculo = Veiculo{Placa: registro[0], Categoria: registro[1], Combustivel: registro[2]}

	var veiculoAsBytes, _ = json.Marshal(veiculo)

	fmt.Print(registro[0])
	fmt.Println(veiculoAsBytes)
	//fmt.Println(json.Unmarshal(veiculoAsBytes, &veiculo))
	fmt.Println(veiculo)
	*/
}
