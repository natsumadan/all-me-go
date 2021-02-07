package main

import (
	"fmt"
)

func main() {

	x := 10.0
	y := 12.0

	fmt.Printf("The value at x is %v and type is %T", x, x)
	fmt.Printf("\nThe value at y is %v and type is %T", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("The mean of x and y is %v and type is %T", mean, mean)
}
