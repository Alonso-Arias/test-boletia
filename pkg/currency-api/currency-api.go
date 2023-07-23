package currencyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/Alonso-Arias/test-boletia/log"
	"github.com/Alonso-Arias/test-boletia/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "context")

var urlBase = "https://api.currencyapi.com/v3/currencies"

var apiKey = "cur_live_mhcdXGJOTpPgfyrnE5WWXxsGAysjzHpzvQJT5HOg"

func FindCurrencies(currency string) (model.CurrencyData, error) {
	log := loggerf.WithField("func", "FindCurrencies")

	params := url.Values{
		"apikey":     {apiKey},
		"currencies": {currency},
	}

	fullURL := urlBase + "?" + params.Encode()

	fmt.Print("BASE ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return model.CurrencyData{}, err
	}
	req.Header.Set("Accept", "application/json")

	client := getClientCofiguration()

	resp, err := client.Do(req)
	if err != nil {
		return model.CurrencyData{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.CurrencyData{}, err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf")) // evitar problemas de formato utf

	var localsResponse model.CurrencyData
	// se parsea el json a la structura declarada
	err = json.Unmarshal(bodyBytes, &localsResponse)
	if err != nil {
		return model.CurrencyData{}, err
	}

	log.Debugf("Body: %s", string(bodyBytes))

	return localsResponse, err

}

// se configuran los deadines de peticiones hacia la api
func getClientCofiguration() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Duration(5) * time.Second,
				KeepAlive: time.Duration(5),
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	return client
}
