package main

import (
	"fmt"

	"github.com/narumiruna/go-visa-fx-rates/pkg/visa"
)

func main() {
	usdtwd, err := visa.Rate("USD", "TWD")
	if err != nil {
		panic(err)
	}
	fmt.Println(usdtwd)
}
