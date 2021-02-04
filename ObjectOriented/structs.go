package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int64
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 10, 12.32, true}
	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Println(t3)
}
