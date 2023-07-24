package dao

import (
	"testing"
	"time"

	"github.com/Alonso-Arias/test-boletia/db/model"
	"github.com/stretchr/testify/assert"
)

func TestCallsLogDAO_Save(t *testing.T) {
	// Crear una instancia del DAO
	callsLogDAO := NewCallsLogDAO()

	// Crear un registro ficticio para guardar
	callsLog := model.CallsLog{
		CallTimestamp:  time.Now(),
		ResponseTimeMs: 654,
		Status:         "Success",
	}

	// Guardar el registro en la base de datos
	err := callsLogDAO.Save(callsLog)
	assert.NoError(t, err, "Error al guardar el registro")
}
