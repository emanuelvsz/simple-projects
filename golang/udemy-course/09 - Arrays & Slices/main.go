package main

import "fmt"

func main() {
	var arr [5]int

	arr[0] = 4
	arr[1] = 3
	arr[2] = 2
	arr[3] = 5
	arr[4] = 1

	fmt.Println(arr)

	arr2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr2)

	slice := []int{
		1, 2, 3,
	}
	slice = append(slice, 100)

	fmt.Println(slice)
}
