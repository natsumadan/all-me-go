package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {

	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {

	t := Trade{
		Volume: 25,
		Symbol: "MSFT",
		Price:  45.36,
		Buy:    true,
	}

	fmt.Println(t.Value())

}
