package main

type Veiculo struct {
	Categoria string `json:"Categoria"`
	Marca     string `json:"Marca"`
	Versao    string `json:"Versao"`
	Modelo    string `json:"Modelo"`
	Emissao   int    `json:"Emissao"`
	Placa     string `json:"Placa"`
}

func main() {
	veiculoInfo := Veiculo{Categoria: "Compacto", Marca: "VW", "Versao": "GOL", "Modelo': 'PATRULHEIRO', 'Emissao': 113.08668705501557}
}
