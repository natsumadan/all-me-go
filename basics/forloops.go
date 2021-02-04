package main

import (
	"fmt"
)

func main() {

	fmt.Println("----")
	for i := 0; i < 15; i++ {
		if i > 4 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	a := 0
	for a < 3 {
		fmt.Println(a)

	}
}
