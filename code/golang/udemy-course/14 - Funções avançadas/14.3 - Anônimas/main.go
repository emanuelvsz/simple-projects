package main

import "fmt"

func main() {
	func(name string) {
		fmt.Println("Hello World! Hi", name, "!")
	}("Emanuel")
}
