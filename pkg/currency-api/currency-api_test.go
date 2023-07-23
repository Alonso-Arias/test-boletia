package currencyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCurrencies_OK(t *testing.T) {
	// Configuramos el API Key y la moneda para la prueba
	currency := "USD" // Cambia esto a la moneda que esperas recibir

	// Ejecutamos la función bajo prueba
	res, err := FindCurrencies(currency)

	// Afirmamos que no hay error
	assert.NoError(t, err, "FindCurrencies error")

	// Afirmamos que el resultado no es nulo
	assert.NotNil(t, res, "Null result")

	// Afirmamos que la respuesta contiene datos válidos
	for _, v := range res.Data {
		assert.NotEmpty(t, v.Symbol, "No data for response")
	}

	// Afirmamos que la moneda devuelta es la correcta
	for _, v := range res.Data {
		assert.Equal(t, currency, v.Code, "Incorrect currency")
	}
}
