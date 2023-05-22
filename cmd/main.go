package main

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/narumiruna/go-visa-fx-rates/pkg/visa"
)

func main() {
	c := visa.NewRestClient()

	i := 0
	err := retry.Do(
		func() error {
			fmt.Println("retry: #", i)
			i++
			now := time.Now()
			request := visa.RatesRequest{
				Amount:           1.0,
				UTCConvertedDate: now,
				ExchangeDate:     now,
				FromCurr:         "TWD",
				ToCurr:           "USD",
			}
			response, err := c.Rates(context.Background(), request)

			if err != nil {
				yesterday := now.AddDate(0, 0, -1)
				request.UTCConvertedDate = yesterday
				request.ExchangeDate = yesterday
				response, err = c.Rates(context.Background(), request)
				if err != nil {
					return err
				}
			}

			fmt.Println(response)

			return nil
		},
	)

	if err != nil {
		panic(err)
	}
	// usdtwd, err := visa.Rate("TWD", "USD")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(usdtwd)
}
