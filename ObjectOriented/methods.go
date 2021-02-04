package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume float64
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
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}

	fmt.Println(t.Value())
	fmt.Printf("t is %+v", t.Value())

}
