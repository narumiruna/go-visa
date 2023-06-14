package visa

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/avast/retry-go"
)

const defaultTimeout = 5 * time.Second
const baseApiUrl = "http://www.visa.com.tw"

type RestClient struct {
	client  *http.Client
	baseURL *url.URL
}

func NewRestClient() *RestClient {
	u, err := url.Parse(baseApiUrl)
	if err != nil {
		panic(err)
	}

	return &RestClient{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: u,
	}
}

func (c *RestClient) NewRequest(ctx context.Context, method string, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	if params != nil {
		rel.RawQuery = params.Encode()
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	req = req.WithContext(ctx)
	return req, nil
}

// http://www.visa.com.tw/cmsapi/fx/rates?amount=100000&fee=0&utcConvertedDate=05%2F01%2F2023&exchangedate=05%2F01%2F2023&fromCurr=TWD&toCurr=USD
func (c *RestClient) CalculateConversion(ctx context.Context, request RatesRequest) (response *RatesResponse, err error) {
	req, err := c.NewRequest(ctx, "GET", "/cmsapi/fx/rates", request.Values())
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *RestClient) ExchangeRate(ctx context.Context, from, to string) (v float64, err error) {
	err = retry.Do(
		func() error {
			now := time.Now()
			request := RatesRequest{
				Amount:           1.0,
				UTCConvertedDate: now,
				ExchangeDate:     now,
				FromCurr:         from,
				ToCurr:           to,
			}
			response, err := c.CalculateConversion(ctx, request)
			if err != nil {
				yesterday := now.AddDate(0, 0, -1)
				request.UTCConvertedDate = yesterday
				request.ExchangeDate = yesterday
				response, err = c.CalculateConversion(ctx, request)
				if err != nil {
					return err
				}
			}

			v, err = strconv.ParseFloat(response.ConvertedAmount, 64)
			return err
		},
	)
	return v, err
}

func (c *RestClient) AskPrice(ctx context.Context, baseCurrency string, quoteCurrency string) (float64, error) {
	return c.ExchangeRate(ctx, quoteCurrency, baseCurrency)
}

func (c *RestClient) BidPrice(ctx context.Context, baseCurrency string, quoteCurrency string) (float64, error) {
	v, err := c.ExchangeRate(ctx, baseCurrency, quoteCurrency)
	if err != nil {
		return 0, err
	}
	return 1 / v, nil
}
