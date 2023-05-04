package questions

import "fmt"

func ListPositives() {
	numList := []int{1, 2, 3, 4, 5, 5, 6, 7, 73, 4, -3, 4, 4, 4, -9, 2}

	for _, n := range numList {
		if n > 0 {
			fmt.Println("O número ", n, "é positivo")
		} else {
			fmt.Println("O número ", n, "é negativo. Parou!")
			break
		}
	}
}
