package main

import (
	"fmt"
	"learn/src/core/domain"

	"github.com/google/uuid"
)

func main() {
	user1 := domain.NewBuilder()
	user1.WithID(uuid.New()).WithName("Emanuel").WithAge(19)
	user1.Build()

	
	
	fmt.Println(&user1)
}
