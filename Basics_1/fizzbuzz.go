package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 51; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
