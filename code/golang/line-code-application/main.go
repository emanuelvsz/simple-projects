package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Init application")

	application := app.Gerar()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
