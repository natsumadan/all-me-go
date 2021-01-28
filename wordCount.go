package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	words := map[string]int{}
	// text2 := map [strings.Fields] {}

	secondHello := strings.ToLower(text)
	fmt.Printf("strings are: %q", strings.Fields(secondHello))
	hello := strings.Fields(secondHello)

	//words := map[string]int{}
	for _, chello := range hello {
		fmt.Println(chello)
		value, ok := words[chello]
		if ok {
			words[chello] = value + 1
		} else {
			words[chello] = 1
		}

	}
	fmt.Println(words)

}
