package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v {type  %T} \n", book[0], book[0])

	//strings are immutabel
	fmt.Println(book[4:11])

	//slice  (no end )

	fmt.Println(book[4:])

	//slice (no strat)
	fmt.Println(book[:4])

	// USe + to conncerate strings
	fmt.Println("t" + )
}
