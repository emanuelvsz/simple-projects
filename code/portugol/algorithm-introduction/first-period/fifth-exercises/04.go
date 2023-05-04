package questions

import "fmt"

type Comodo struct {
	Nome        string
	Largula     float64
	Comprimento float64
}

func NewComodo(nome string, largura, comprimento float64) *Comodo {
	return &Comodo{
		Nome:        nome,
		Largula:     largura,
		Comprimento: comprimento,
	}
}

func CalcRooms() {
	count := 0
	comodos := make([]Comodo, count)

	for i := -1; i < count; i++ {
		var nome string
		var largura float64
		var comprimento float64

		fmt.Println("Digite um nome ao cômodo")
		fmt.Scan(&nome)
		fmt.Println("Digite uma largura")
		fmt.Scan(&largura)
		fmt.Println("Digite um comprimento")
		fmt.Scan(&comprimento)

		var chs bool
		fmt.Println("\n\nDeseja adicionar mais um cômodo?(true ou false)")
		fmt.Scan(&chs)

		newComodos := NewComodo(nome, largura, comprimento)
		comodos = append(comodos, *newComodos)

		if chs {
			count++
		} else {
			fmt.Println(comodos)
		}
	}
}