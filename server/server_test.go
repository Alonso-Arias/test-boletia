package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alonso-Arias/test-boletia/handler"
	"github.com/Alonso-Arias/test-boletia/services/currency"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCurrenciesGet(t *testing.T) {
	// Crea un contexto de prueba
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/currencies/USD", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Llama al handler CurrenciesGet
	err := handler.CurrenciesGet(c)

	// Verifica si no hay errores
	assert.NoError(t, err)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verifica el contenido de la respuesta
	var result currency.FindCurrenciesResponse
	err = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)

	// Verifica si el resultado coincide con lo esperado
	expectedMessage := currency.FindCurrenciesResponse{
		// Agrega aquí los campos esperados en la respuesta, según la estructura de currency.FindCurrenciesResponse
	}
	assert.Equal(t, expectedMessage, result)
}
