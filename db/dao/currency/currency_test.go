package dao

import (
	"context"
	"testing"
	"time"

	"github.com/Alonso-Arias/test-boletia/db/model"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyDAO_Save(t *testing.T) {
	// Crear una instancia del DAO
	currencyDAO := NewCurrencyDAO()

	// Crear una Currency ficticia para guardarla
	currency := model.Currency{
		Currency:  "USD",
		Value:     1.23,
		Timestamp: time.Now(),
	}

	// Guardar la Currency en la base de datos
	err := currencyDAO.Save(currency)
	assert.NoError(t, err, "Error al guardar la Currency")
}

func TestCurrencyDAO_FindCurrencyValuesByDate(t *testing.T) {
	// Crear una instancia del DAO
	currencyDAO := NewCurrencyDAO()

	// Definir fechas de inicio y fin para la búsqueda
	start := time.Now().AddDate(0, 0, -7) // Hace 7 días
	end := time.Now()

	// Definir la moneda a buscar
	currency := "USD"

	// Realizar la búsqueda de Currency por fecha y moneda
	currencies, err := currencyDAO.FindCurrencyValuesByDate(context.Background(), start, end, currency)
	assert.NoError(t, err, "Error al buscar Currency por fecha y moneda")
	assert.NotEmpty(t, currencies, "No se encontraron Currency para las fechas y moneda dadas")
}

func TestCurrencyDAO_FindCurrencyValueDate(t *testing.T) {
	// Crear una instancia del DAO
	currencyDAO := NewCurrencyDAO()

	// Definir la moneda y el tipo de fecha (FirstDate o LastDate)
	currency := "USD"
	dateType := "FirstDate"

	// Realizar la búsqueda de la fecha más antigua para la Currency dada
	date, err := currencyDAO.FindCurrencyValueDate(context.Background(), currency, dateType)
	assert.NoError(t, err, "Error al buscar la fecha de Currency")
	assert.NotEqual(t, time.Time{}, date, "No se encontró ninguna fecha para la Currency dada")
}

func TestCurrencyDAO_GetAllCurrenciesWithLatestValue(t *testing.T) {
	// Crear una instancia del DAO
	currencyDAO := NewCurrencyDAO()

	// Obtener todas las monedas con su valor más reciente
	currencies, err := currencyDAO.GetAllCurrenciesWithLatestValue(context.Background())
	assert.NoError(t, err, "Error al obtener todas las monedas con el valor más reciente")
	assert.NotEmpty(t, currencies, "No se encontraron monedas con valores más recientes")
}
