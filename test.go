package main

import "fmt"

func main() {
	isAvaiable := true

	IsAvaiableValid(isAvaiable)
}

func IsAvaiableValid(isAvaiable bool) bool {
	if isAvaiable || !isAvaiable {
		fmt.Println("É valido boy")
		return true
	}

	fmt.Println("É invalido girl")
	return false
}
