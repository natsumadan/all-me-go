package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("X is big ")
	}
	if x > 100 {
		fmt.Println("X is not that bigfv ")
	} else {
		fmt.Println("X is not tghat n=big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is sammler than ")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more thatn half of b")
	}

}
