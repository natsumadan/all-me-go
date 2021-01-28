package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := 0

	for i, he := range nums {
		if max < nums[i] {
			max = nums[i]
		}

	}
	fmt.Println(max)
}
