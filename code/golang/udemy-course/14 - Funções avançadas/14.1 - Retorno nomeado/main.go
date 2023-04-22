package main

func calcMatematicos(num1, num2 int) (soma, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

func main() {
	calcMatematicos(1, 2)
}
