package questions

import "fmt"

func HundredSum() {
	var nums []uint8

	var sum uint32
	for i := 0; i < 101; i++ {
		nums = append(nums, uint8(i))
	}

	for _, n := range nums {
		sum += uint32(n)
	}

	fmt.Println(sum)
}
