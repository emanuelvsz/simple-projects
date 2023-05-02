package main

import (
	"bufio"
	"fmt"
	"morse/functions"
	"os"
)

func main() {
	morse := "./--/.-/-./..-/./.-.."

	fmt.Println("Digite uma frase para ser convertida em código morse: \n")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')

	fmt.Println("Resultado: \n\n")

	fmt.Println("Frase escrita:", frase)
	print(morse, frase)
}

func print(morse, phrase string) {
	fmt.Println("Frase criptografada:", functions.ConvertStringToMorse(phrase), "\n")
	fmt.Println("Morse descriptografado:", functions.ConvertMorseToString(morse))
}
