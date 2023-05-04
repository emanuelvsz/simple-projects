package questions

import (
	"fmt"
	"sort"
)

func ListNumbers() {
	nums := []int{0, 10}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var newNums []int
	for i := nums[1]; i <= nums[0]; i++ {
		if i%2 == 0 {
			newNums = append(newNums, i)
		}
	}

	fmt.Println(newNums)
	fmt.Println(nums)
}
