package visa

import (
	"fmt"
	"net/url"
	"time"
)

type RateOptions struct {
	Amount           float64   `json:"amount"`
	Fee              float64   `json:"fee"`
	UTCConvertedDate time.Time `json:"utcConvertedDate"`
	ExchangeDate     time.Time `json:"exchangedate"`
	FromCurr         string    `json:"fromCurr"`
	ToCurr           string    `json:"toCurr"`
}

func (o *RateOptions) Values() url.Values {
	values := url.Values{}

	values.Add("amount", fmt.Sprintf("%f", o.Amount))
	values.Add("fee", fmt.Sprintf("%f", o.Fee))
	values.Add("utcConvertedDate", o.UTCConvertedDate.Format("01/02/2006"))
	values.Add("exchangedate", o.ExchangeDate.Format("01/02/2006"))
	values.Add("fromCurr", o.FromCurr)
	values.Add("toCurr", o.ToCurr)

	return values
}

//	{
//		"originalValues":{
//		   "fromCurrency":"USD",
//		   "fromCurrencyName":"United States Dollar",
//		   "toCurrency":"TWD",
//		   "toCurrencyName":"New Taiwan Dollar",
//		   "asOfDate":1682899200,
//		   "fromAmount":"100000",
//		   "toAmountWithVisaRate":"3074900",
//		   "toAmountWithAdditionalFee":"3074900",
//		   "fxRateVisa":"30.749",
//		   "fxRateWithAdditionalFee":"30.749",
//		   "lastUpdatedVisaRate":1682725860,
//		   "benchmarks":[]
//		},
//		"conversionAmountValue":"100000",
//		"conversionBankFee":"0.0",
//		"conversionInputDate":"05/01/2023",
//		"conversionFromCurrency":"TWD",
//		"conversionToCurrency":"USD",
//		"fromCurrencyName":"United States Dollar",
//		"toCurrencyName":"New Taiwan Dollar",
//		"convertedAmount":"3,074,900.000000",
//		"benchMarkAmount":"",
//		"fxRateWithAdditionalFee":"30.749",
//		"reverseAmount":"0.032521",
//		"disclaimerDate":"May 1, 2023",
//		"status":"success"
//	}
type RateResponse struct {
	OriginalValues          OriginalValues `json:"originalValues"`
	ConversionAmountValue   string         `json:"conversionAmountValue"`
	ConversionBankFee       string         `json:"conversionBankFee"`
	ConversionInputDate     string         `json:"conversionInputDate"`
	ConversionFromCurrency  string         `json:"conversionFromCurrency"`
	ConversionToCurrency    string         `json:"conversionToCurrency"`
	FromCurrencyName        string         `json:"fromCurrencyName"`
	ToCurrencyName          string         `json:"toCurrencyName"`
	ConvertedAmount         string         `json:"convertedAmount"`
	BenchMarkAmount         string         `json:"benchMarkAmount"`
	FxRateWithAdditionalFee string         `json:"fxRateWithAdditionalFee"`
	ReverseAmount           string         `json:"reverseAmount"`
	DisclaimerDate          string         `json:"disclaimerDate"`
	Status                  string         `json:"status"`
}

type OriginalValues struct {
	FromCurrency              string        `json:"fromCurrency"`
	FromCurrencyName          string        `json:"fromCurrencyName"`
	ToCurrency                string        `json:"toCurrency"`
	ToCurrencyName            string        `json:"toCurrencyName"`
	AsOfDate                  int           `json:"asOfDate"`
	FromAmount                string        `json:"fromAmount"`
	ToAmountWithVisaRate      string        `json:"toAmountWithVisaRate"`
	ToAmountWithAdditionalFee string        `json:"toAmountWithAdditionalFee"`
	FxRateVisa                string        `json:"fxRateVisa"`
	FxRateWithAdditionalFee   string        `json:"fxRateWithAdditionalFee"`
	LastUpdatedVisaRate       int           `json:"lastUpdatedVisaRate"`
	Benchmarks                []interface{} `json:"benchmarks"`
}
