package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "Cidade A", Destino: "Cidade B", Data: time.Now(), Preco: 500.0},
		{Origem: "Cidade X", Destino: "Cidade Y", Data: time.Now(), Preco: 800.0},
		{Origem: "Cidade M", Destino: "Cidade N", Data: time.Now(), Preco: 600.0},
	}

	viagemMaisCara := viagemMaisCara(viagens)

	fmt.Printf("A viagem mais cara é de %s para %s, com preço de %.2f\n", viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Preco)
}
