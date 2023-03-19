package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 5 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	users := [3]string{"Emanuel", "Pedro", "Leo"}

	for _, name := range users {
		fmt.Println(name)
	}
}
