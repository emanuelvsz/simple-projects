package main

import "fmt"

type Person struct {
	ID   int64
	Name string
	Age  uint8
}

type Studant struct {
	Person
	Course string
}

func main() {

	studant := Studant{
		Person: Person{
			ID:   123,
			Name: "Emanuel Vilela",
			Age:  19,
		},
		Course: "Sistemas de Informação",
	}

	fmt.Println(studant)
}
