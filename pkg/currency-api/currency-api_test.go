package currencyapi

import (
	"testing"

	"github.com/Alonso-Arias/test-boletia/pkg/currency-api/client"
	"github.com/stretchr/testify/assert"
)

func TestFindCurrencies_OK(t *testing.T) {
	// Configuramos el API Key y la moneda para la prueba

	// Ejecutamos la función bajo prueba
	res, _, err := client.FindCurrencies()

	// Afirmamos que no hay error
	assert.NoError(t, err, "FindCurrencies error")

	// Afirmamos que el resultado no es nulo
	assert.NotNil(t, res, "Null result")

	// Afirmamos que la respuesta contiene datos válidos
	for _, v := range res.Data {
		assert.NotEmpty(t, v.Value, "No data for response")
	}
}
