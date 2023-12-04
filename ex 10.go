package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) calcularMediaAvaliacoes() float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}

	return float64(soma) / float64(len(f.Avaliacoes))
}

func (f *Filme) imprimirInfo() {
	fmt.Printf("Título: %s\nDiretor: %s\nAno: %d\nMédia de Avaliações: %.2f\n", f.Titulo, f.Diretor, f.Ano, f.calcularMediaAvaliacoes())
}

func main() {
	filmeExemplo := Filme{
		Titulo:     "O Filme",
		Diretor:    "Diretor A",
		Ano:        2022,
		Avaliacoes: []int{4, 5, 3, 4, 5},
	}

	filmeExemplo.imprimirInfo()

	filmeExemplo.adicionarAvaliacao(5)
	fmt.Println("\nAdicionando nova avaliação 5:")
	filmeExemplo.imprimirInfo()

	filmeExemplo.removerAvaliacao(2)
	fmt.Println("\nRemovendo a terceira avaliação:")
	filmeExemplo.imprimirInfo()
}
