package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 12, 25, 3, 64, 89, 20, 10}

	//
	for i := range nums {
		for j := range nums {
			if i == j {
				fmt.Println("true")

			}

		}

	}

}
