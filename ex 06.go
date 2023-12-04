package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) aumentarSalario(porcentagem float64) {
	f.Salario *= (1 + porcentagem/100)
}

func (f *Funcionario) diminuirSalario(porcentagem float64) {
	f.Salario *= (1 - porcentagem/100)
}

func (f *Funcionario) tempoDeServico() time.Duration {
	idadeInicioTrabalho := 18
	anosDeServico := f.Idade - idadeInicioTrabalho
	return time.Duration(anosDeServico) * 365 * 24 * time.Hour
}

func main() {
	funcionarioExemplo := Funcionario{
		Nome:    "João",
		Salario: 5000.0,
		Idade:   30,
	}

	fmt.Printf("Salário antes do aumento: %.2f\n", funcionarioExemplo.Salario)
	funcionarioExemplo.aumentarSalario(10)
	fmt.Printf("Salário depois do aumento: %.2f\n", funcionarioExemplo.Salario)

	fmt.Printf("Salário antes da redução: %.2f\n", funcionarioExemplo.Salario)
	funcionarioExemplo.diminuirSalario(5)
	fmt.Printf("Salário depois da redução: %.2f\n", funcionarioExemplo.Salario)

	tempoServico := funcionarioExemplo.tempoDeServico()
	fmt.Printf("Tempo de serviço do funcionário: %v\n", tempoServico)
}
