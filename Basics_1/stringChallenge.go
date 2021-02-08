package main

import (
	"fmt"
)

func main() {

	count := 0
	for i := 1000; i < 9999; i++ {
		for j := 1000; j < 9999; j++ {
			n := i * j

			//fmt.Println(n)

			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
				//fmt.Println(s)
			}

		}
	}

	fmt.Println(count)

}
