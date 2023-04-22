package pacotes

import (
	"fmt"
)

// Write "escreve no terminal uma frase em inglês"
func Write() {
	fmt.Println("Writing in the package 'pacotes'")
	write()
}

func write() {
	fmt.Println("Writing in the package 'pacotes' with the private method")
}
