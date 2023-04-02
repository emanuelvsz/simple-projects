package main

import (
	"fmt"
	"learn/src/core/domain"

	"github.com/google/uuid"
)

func main() {
	user1 := domain.NewBuilder()
	user1.WithID(uuid.New()).WithName("Emanuel").WithAge(19)
	emanuel, err := user1.Build()
	if err != nil {
		fmt.Println("houve um erro na criação do domain")
	}

	user2 := domain.NewBuilder()
	user2.WithID(uuid.New()).WithName("Emanuel").WithAge(19)
	douglas, err := user1.Build()
	if err != nil {
		fmt.Println("houve um erro na criação do domain")
	}

	alunos := []domain.Aluno{*emanuel, *douglas}

	fmt.Println(alunos)
}
