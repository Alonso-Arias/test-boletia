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

	client := getClientConfiguration()

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
// getClientConfiguration configura los timeouts y deadlines de las peticiones hacia la API
func getClientConfiguration() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second, // Tiempo máximo para establecer la conexión
				KeepAlive: 5 * time.Second, // Tiempo de vida de la conexión abierta
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second, // Tiempo máximo para completar el handshake TLS
			ResponseHeaderTimeout: 5 * time.Second, // Tiempo máximo para recibir la respuesta después del handshake
			ExpectContinueTimeout: 1 * time.Second, // Tiempo máximo para recibir una respuesta después de enviar "Expect: 100-continue"
		},
		Timeout: 15 * time.Second, // Tiempo máximo para realizar una operación completa, incluyendo conexión, handshake y respuesta del servidor
	}

	return client
}
