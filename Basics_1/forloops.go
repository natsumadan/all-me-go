package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------")
	for i := 0; i < 3; i++ {
		if i > 1 {
			fmt.Println(i)
		}
	}
}
