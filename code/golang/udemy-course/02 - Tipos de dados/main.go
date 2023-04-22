package main

import (
	"errors"
	"fmt"
)

func main() {
	positiveOrNegativeIntegers()
	justPositiveIntegers()
	crackedNumbers()
	textValues()
	boolValues()
	errorValues()
}

func positiveOrNegativeIntegers() {
	fmt.Println("Valores inteiros(int, rune, byte)")

	var number int8 = 100
	fmt.Println(number)

	var number2 int16 = 10000
	fmt.Println(number2)

	var number3 int32 = 1000000000
	fmt.Println(number3)

	var number4 int64 = -1000000000000000000
	fmt.Println(number4)

	var number5 int = 1010101011000
	fmt.Println(number5)

	var number6 rune = 12323
	fmt.Println(number6) // um alias pro int32

	var number7 byte = 112
	fmt.Println(number7) // um alias pro uint8
}

func crackedNumbers() {
	fmt.Println("Valores floaters(float32 e float64)")

	var number float32 = 100.1231231231
	fmt.Println(number)

	var number2 float64 = 10000.1123123123
	fmt.Println(number2)
}

func justPositiveIntegers() {
	fmt.Println("Valores inteiros somente positivos(uint)")

	var number uint8 = 100
	fmt.Println(number)
}

func textValues() {
	fmt.Println("Valores de texto(string)")

	var nome string = "Emanuel"
	fmt.Println(nome)

	nome2 := "E"
	fmt.Println(nome2)

}

func boolValues() {
	fmt.Println("Valores de decisão(bool)")

	var dcd bool
	fmt.Println(dcd) // por padrão, retorna false

	var dcd2 bool = true
	fmt.Println(dcd2)
}

func errorValues() {
	fmt.Println("Valores de erro(error)")

	var erro error // por padrão retorna nil
	fmt.Println(erro)

	var erro2 error = errors.New("Erro de tal tal tal")
	fmt.Println(erro2)
}
