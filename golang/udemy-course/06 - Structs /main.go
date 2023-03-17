package main

import "fmt"

func main() {
	fmt.Println("Aprendendo sobre structs")

	users := []User{
		{
			Nome:  "Emanuel Vilela",
			Idade: 19,
		},
		{
			Nome:  "Pedro Vilela",
			Idade: 20,
		},
	}
	var user User
	fmt.Println(users)

	user.Idade = 22
	user.Nome = "Emanuel"

	fmt.Println(user)

}

type User struct {
	Nome  string
	Idade int
}
