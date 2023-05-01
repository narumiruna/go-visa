package main

import (
	"context"
	"fmt"
	"time"

	"github.com/narumiruna/go-visa-fx-rates/pkg/visa"
)

func main() {
	c := visa.NewRestClient()
	now := time.Now()
	options := visa.RatesRequest{
		Amount:           1,
		FromCurr:         "USD",
		ToCurr:           "TWD",
		Fee:              0.0,
		UTCConvertedDate: now,
		ExchangeDate:     now,
	}
	resp, err := c.Rates(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %+v\n", resp)
}
