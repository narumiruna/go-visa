package visa

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 5 * time.Second
const baseApiUrl = "https://www.visa.com.tw"

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

	req = req.WithContext(ctx)
	return req, nil
}

// https://www.visa.com.tw/cmsapi/fx/rates?amount=100000&fee=0&utcConvertedDate=05%2F01%2F2023&exchangedate=05%2F01%2F2023&fromCurr=TWD&toCurr=USD
func (c *RestClient) Rates(ctx context.Context, request RatesRequest) (response *RatesResponse, err error) {
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
