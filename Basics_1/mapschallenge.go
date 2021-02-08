package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	words := strings.Fields(text)
	//ss:= strings.ToLower(s)

	co := map[string]int{}
	for _, word := range words {
		co[strings.ToLower(word)]++
	}

	fmt.Println(co)

}
