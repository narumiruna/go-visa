package visa

import "context"

func ExchangeRate(from, to string) (v float64, err error) {
	c := NewRestClient()
	r, err := c.ExchangeRate(context.Background(), from, to)
	if err != nil {
		return 0, err
	}
	return r, nil
}
