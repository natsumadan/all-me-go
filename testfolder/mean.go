//calculate the mean of the two numbers

package main

import (
	"fmt"
)

func main() {
	//var x float64
	//var y float64

	x := 1.0
	y := 2.0

	fmt.Printf("x=%v , type of %T\n", x, y)
	fmt.Printf("y=%v , type of %T\n", y, y)

	measn := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", measn, measn)

}
