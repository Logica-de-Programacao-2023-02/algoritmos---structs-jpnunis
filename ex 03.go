package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func calcularAreaTriangulo(t Triangulo) float64 {
	return (t.Base * t.Altura) / 2
}

func main() {
	trianguloExemplo := Triangulo{
		Base:   4.0,
		Altura: 6.0,
	}

	area := calcularAreaTriangulo(trianguloExemplo)
	fmt.Printf("A área do triângulo com base %.2f e altura %.2f é %.2f\n", trianguloExemplo.Base, trianguloExemplo.Altura, area)
}
