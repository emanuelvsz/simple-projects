package main

import "fmt"

func soma(numeros ...int) {
	for _, val := range numeros {
		fmt.Println(val)
	}
}

func main() {
	soma(1, 2, 3, 4, 5)
}
