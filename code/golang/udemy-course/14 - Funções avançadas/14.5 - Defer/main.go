package main

import "fmt"

func funcUm() {
	fmt.Println("Opa1!")
}

func funcDois() {
	fmt.Println("Opa2!")
}

func main() {
	defer funcUm()
	funcDois()
}
