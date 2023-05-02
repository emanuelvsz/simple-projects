package main

import "fmt"

type ContaCorrente struct {
	dono    string
	saldo   float64
	agencia string
}

func (instance ContaCorrente) Sacar() string {
	
	return "Saldo insuficiente"
}

func main() {
	var contas []ContaCorrente

	conta1 := ContaCorrente{}
	conta1.agencia = "Caixa"
	conta1.dono = "Emanuel"
	conta1.saldo = 1000

	contas = append(contas, conta1)

	fmt.Println(contas)
}
