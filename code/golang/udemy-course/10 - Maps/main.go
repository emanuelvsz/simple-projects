package main

import "fmt"

func main() {
	user := []map[string]string{
		{
			"name": "Emanuel",
			"age":  "",
		},
		{
			"name": "Diogo",
			"age":  "21",
		},
	}
	//delete(user[1], "age")

	for i := 0; i < len(user); i++ {
		if len(user[i]["age"]) == 0 {
			fmt.Println("Usuário", user[i]["name"], "sem idade cadastrada no sistema.")
		} else {
			fmt.Println(user[i])
		}
	}
}
