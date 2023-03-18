package main

import "fmt"

func main() {
	fmt.Println("Ponteiros: ")

	var name string = "Emanuel"
	var ponteiro *string = &name

	fmt.Println(*ponteiro)
}
