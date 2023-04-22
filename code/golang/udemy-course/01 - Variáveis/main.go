package main

import "fmt"

func main() {
	var name string = "Nome com tipagem"      // string com declaração de tipo
	nameTip := "Nome sem tipagem determinada" // string declarada sem tipagem

	fmt.Println(name)
	fmt.Println(nameTip)

	var (
		var1 int = 1
		var2 int = 2
	)

	fmt.Println(var1)
	fmt.Println(var2)

	var3, var4 := 3, 4

	fmt.Println(var3)
	fmt.Println(var4)
}
