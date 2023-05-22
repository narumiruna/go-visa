package visa

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	log "github.com/sirupsen/logrus"
)

func rate(request RatesRequest) (response *RatesResponse, err error) {
	u, err := url.Parse("https://www.visa.com.tw/cmsapi/fx/rates")
	if err != nil {
		return nil, err
	}

	u.RawQuery = request.Values().Encode()
	log.Infof("url: %s", u.String())

	b := rod.New()
	defer b.MustClose()

	// connect to browser
	err = b.Connect()
	if err != nil {
		return nil, err
	}

	// set timeout
	b = b.Timeout(defaultTimeout)

	page, err := b.Page(proto.TargetCreateTarget{URL: u.String()})
	if err != nil {
		return nil, err
	}

	element, err := page.Element("html")
	if err != nil {
		return nil, err
	}

	text, err := element.Text()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(text), &response); err != nil {
		return nil, err
	}

	return response, nil
}

func Rate(from, to string) (float64, error) {
	now := time.Now()
	request := RatesRequest{
		Amount:           1.0,
		Fee:              0.0,
		UTCConvertedDate: now,
		ExchangeDate:     now,
		FromCurr:         from,
		ToCurr:           to,
	}

	response, err := rate(request)
	if err != nil {
		// try yesterday
		yesterday := now.AddDate(0, 0, -1)
		request.ExchangeDate = yesterday
		request.UTCConvertedDate = yesterday

		response, err = rate(request)
		if err != nil {
			return 0, err
		}
	}

	return strconv.ParseFloat(response.FxRateWithAdditionalFee, 64)
}
