package main

import (
	"fmt"
	pacotes "modulo/packages"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing in the code.")
	pacotes.Write()

	err := checkmail.ValidateFormat("usermane13@gmail.com")
	fmt.Println(err)

	if err != nil {
		fmt.Println("O texto escrito não é um email")
	}

}
