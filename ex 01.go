package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	circuloExemplo := Circulo{raio: 3.0}
	area := calcularArea(circuloExemplo)
	fmt.Printf("A área do círculo com raio %.2f é %.2f\n", circuloExemplo.raio, area)
}
