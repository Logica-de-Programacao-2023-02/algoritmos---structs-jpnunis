package main

import "fmt"

type Animal struct {
	Nome   string
	Especie string
	Idade  int
	Som    string
}

func (a *Animal) modificarSom(novoSom string) {
	a.Som = novoSom
}

func (a *Animal) imprimirInfo() {
	fmt.Printf("Informações do animal:\nNome: %s\nEspécie: %s\nIdade: %d anos\nSom: %s\n", a.Nome, a.Especie, a.Idade, a.Som)
}

func main() {
	animalExemplo := Animal{
		Nome:    "Rex",
		Especie: "Cachorro",
		Idade:   5,
		Som:     "Latido",
	}

	animalExemplo.imprimirInfo()

	fmt.Println("\nModificando o som do animal:")
	animalExemplo.modificarSom("Rugido")
	animalExemplo.imprimirInfo()
}
