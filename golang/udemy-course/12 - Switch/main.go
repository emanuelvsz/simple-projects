package main

import (
	"fmt"
)

func main() {
	dia := 0

	switch dia {
	case 1:
		fmt.Println("Domingo!")
		fallthrough
	case 2:
		fmt.Println("Segunda!")
	case 3:
		fmt.Println("Terça!")
	case 4:
		fmt.Println("Quarta!")
	case 5:
		fmt.Println("Quinta!")
	case 6:
		fmt.Println("Sexta!")
	case 7:
		fmt.Println("Sábado!")
	default:
		fmt.Println("Número inválido")
	}
}
