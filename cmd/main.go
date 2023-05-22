package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-visa-fx-rates/pkg/visa"
)

func main() {
	c := visa.NewRestClient()
	r, err := c.ExchangeRate(context.Background(), "TWD", "USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
