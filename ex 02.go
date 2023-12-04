package main

import "fmt"

type Endereco struct {
	Rua     string
	Numero  int
	Cidade  string
	Estado  string
}

type Pessoa struct {
	Nome    string
	Idade   int
	Endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço:\nRua: %s\nNúmero: %d\nCidade: %s\nEstado: %s\n", p.Endereco.Rua, p.Endereco.Numero, p.Endereco.Cidade, p.Endereco.Estado)
}

func main() {
	pessoaExemplo := Pessoa{
		Nome:    "João",
		Idade:   25,
		Endereco: Endereco{
			Rua:     "Rua ABC",
			Numero:  123,
			Cidade:  "Cidade XYZ",
			Estado:  "UF",
		},
	}

	imprimirEnderecoCompleto(pessoaExemplo)
}
