package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 16999.64,
		"GOOG": 123551.32,
		"MSFT": 85462.32,
	}

	fmt.Println(stocks["TESLA"])

	value, ok := stocks["GOOG"]
	if !ok {
		fmt.Println("TESLA is not found")
	} else {
		fmt.Println(value)
	}

	stocks["TESLA"] = 582.32
	fmt.Println(stocks)

}
