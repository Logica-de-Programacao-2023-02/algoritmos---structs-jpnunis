package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) removerNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) calcularMedia() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}

	return soma / float64(len(a.Notas))
}

func (a *Aluno) imprimirInfo() {
	fmt.Printf("Nome: %s\nIdade: %d\nMédia das Notas: %.2f\n", a.Nome, a.Idade, a.calcularMedia())
}

func main() {
	alunoExemplo := Aluno{
		Nome:  "Maria",
		Idade: 20,
		Notas: []float64{8.5, 7.0, 9.2},
	}

	alunoExemplo.imprimirInfo()

	alunoExemplo.adicionarNota(9.5)
	fmt.Println("\nAdicionando nova nota 9.5:")
	alunoExemplo.imprimirInfo()

	alunoExemplo.removerNota(1)
	fmt.Println("\nRemovendo a segunda nota:")
	alunoExemplo.imprimirInfo()
}
