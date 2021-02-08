package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, //Must have tailing comma at the end
	}

	// Data structure where keys point to value

	//Number of items
	fmt.Println(len(stocks))

	//Get a value
	fmt.Println(stocks["MSFT"])

	//Get a zero value

	fmt.Println(stocks["TSLA"])

	//Use two zero values to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA is ot found")

	} else {
		fmt.Println(value)
	}

	//Sett a value in stocks
	stocks["TSLA"] = 3220.0
	fmt.Println(stocks)

	//Dekete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	//Single value for is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	//Double value for is key , value
	for key, value := range stocks {

	}

}
