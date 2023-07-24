package currencyapi

import (
	// "time"

	"context"

	daoCl "github.com/Alonso-Arias/test-boletia/db/dao/calls-log"
	daoCu "github.com/Alonso-Arias/test-boletia/db/dao/currency"

	"time"

	"github.com/Alonso-Arias/test-boletia/db/model"

	"github.com/Alonso-Arias/test-boletia/log"
	"github.com/Alonso-Arias/test-boletia/pkg/currency-api/client"
)

var loggerf = log.LoggerJSON().WithField("package", "context")

func FindAndSaveCurrencyValues() {
	log := loggerf.WithField("func", "FindAndSaveCurrencyValues")

	status := "SUCCESS"

	// // se llaman a las instancia de la base de datos
	callsLogDao := daoCl.NewCallsLogDAO()
	currencyDao := daoCu.NewCurrencyDAO()

	// se ejecuta el get a la api
	list, duration, err := client.FindCurrencies()
	if err != context.DeadlineExceeded {
		log.WithError(err).Error("timeout error occurred")
		status = "FAILED"
	}

	// guarda la duracion en milisegundos de la respuesta
	callsLogDao.Save(model.CallsLog{ResponseTimeMs: duration.Milliseconds(), CallTimestamp: time.Now(), Status: status})

	if status == "SUCCESS" {
		// empieza a iterar para guardar todas las divisas con su respectivo valor
		for _, v := range list.Data {
			cv := model.Currency{Currency: v.Code, Value: v.Value, Timestamp: time.Now()}
			err := currencyDao.Save(cv)
			if err != nil {
				log.WithError(err).Error("problems with saving value")
			}
		}
	}

}
