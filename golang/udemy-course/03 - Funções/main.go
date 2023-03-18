package main

import "fmt"

func main() {
	fmt.Println(soma(1, 2))
	varFunction()
	fmt.Println(soma(2, 4))
}

func soma(n1, n2 int) int { // função básica
	return n1 + n2
}

func varFunction() {
	var f = func(txt string) string {
		return txt
	}
	fmt.Println(f("Opa"))
}

func mathMethod(num1, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub
}
