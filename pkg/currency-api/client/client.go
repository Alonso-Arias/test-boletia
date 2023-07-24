package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	errs "github.com/Alonso-Arias/test-boletia/errors"
	"github.com/Alonso-Arias/test-boletia/log"
	"github.com/Alonso-Arias/test-boletia/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "context")

var urlBase = os.Getenv("API_URL")

var apiKey = os.Getenv("API_KEY")

var timeOut = os.Getenv("TIME_OUT_SECONDS")

func FindCurrencies() (model.CurrencyData, time.Duration, error) {
	log := loggerf.WithField("func", "FindCurrencies")

	params := url.Values{
		"apikey": {apiKey},
	}

	// params := url.Values{
	// 	"sleep": {"20000"},
	// }

	fullURL := urlBase + "?" + params.Encode()

	fmt.Print("BASE ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return model.CurrencyData{}, 0, err
	}
	req.Header.Set("Accept", "application/json")

	client := getClientConfiguration()

	// Registro del tiempo de inicio antes de realizar la llamada a la API
	startTime := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return model.CurrencyData{}, 0, err
	}
	defer resp.Body.Close()

	// Cálculo del tiempo de respuesta
	responseTime := time.Since(startTime)

	// Verificación de si el tiempo de respuesta supera el timeout configurado
	err = checkResponseTimeout(responseTime)
	if err != nil {
		return model.CurrencyData{}, responseTime, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.CurrencyData{}, 0, err
	}
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte("\xef\xbb\xbf")) // evitar problemas de formato utf

	var localsResponse model.CurrencyData
	// se parsea el json a la estructura declarada
	err = json.Unmarshal(bodyBytes, &localsResponse)
	if err != nil {
		return model.CurrencyData{}, 0, err
	}

	log.Debugf("Body: %s", string(bodyBytes))

	return localsResponse, responseTime, nil

}

// getClientConfiguration configura los timeouts y deadlines de las peticiones hacia la API
func getClientConfiguration() *http.Client {
	// Convierte el valor de string a int64
	timeOutSeconds, err := strconv.ParseInt(timeOut, 10, 64)
	if err != nil {
		// Manejo de error si no se pudo convertir el valor
		panic(err)
	}
	client := &http.Client{
		Timeout: time.Duration(timeOutSeconds) * time.Second, // Tiempo máximo para realizar una operación completa, incluyendo conexión, handshake y respuesta del servidor
	}

	return client
}

// checkResponseTimeout verifica si el tiempo de respuesta supera el timeout configurado
func checkResponseTimeout(responseTime time.Duration) error {
	// Convierte el valor de string a int64
	timeOutSeconds, err := strconv.ParseInt(timeOut, 10, 64)
	if err != nil {
		// Manejo de error si no se pudo convertir el valor
		return err
	}

	if responseTime > time.Duration(timeOutSeconds)*time.Second {
		// Devuelve un error o un custom error para indicar que se superó el timeout
		return errs.TimeoutErrorApi
	}

	return nil
}
