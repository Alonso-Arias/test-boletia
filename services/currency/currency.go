package currency

import (
	"context"
	"regexp"
	"strings"
	"time"

	daoCu "github.com/Alonso-Arias/test-boletia/db/dao/currency"
	errs "github.com/Alonso-Arias/test-boletia/errors"
	"github.com/Alonso-Arias/test-boletia/log"
	"github.com/Alonso-Arias/test-boletia/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

type CurrencyService struct {
}

type FindCurrenciesRequest struct {
	Currency string    `json:"currency"`
	Finit    time.Time `json:"finit"`
	Fend     time.Time `json:"fend"`
}
type FindCurrenciesResponse struct {
	Currencies map[string][]model.CurrencyResponse `json:"currencies"`
}

func (cs CurrencyService) FindCurrencies(ctx context.Context, in FindCurrenciesRequest) (FindCurrenciesResponse, error) {
	log := loggerf.WithField("service", "CurrencyService").WithField("func", "FindCurrencies")

	log.Info("start")
	defer log.Info("finish")

	// se valida request de entrada
	err := requestValidation(in)
	if err != nil {
		return FindCurrenciesResponse{}, err
	}

	// se valida las entradas de fechas
	request, err := dateValidation(ctx, in)
	if err != nil {
		return FindCurrenciesResponse{}, err
	}

	currencyDao := daoCu.NewCurrencyDAO()

	// Caso especial cuando in.Currency es "all"
	if strings.ToUpper(in.Currency) == "ALL" {
		// Obtener todas las divisas con sus valores más recientes
		allCurrencies, err := currencyDao.GetAllCurrenciesWithLatestValue(ctx)
		if err != nil {
			return FindCurrenciesResponse{}, err
		}

		results := make(map[string][]model.CurrencyResponse)
		for _, currency := range allCurrencies {
			if in.Currency == "all" || currency.Currency == in.Currency {
				data := model.CurrencyResponse{Value: currency.Value, Date: currency.Timestamp.Format("2006-01-02T15:04:05")}
				results[currency.Currency] = append(results[currency.Currency], data)
			}
		}

		return FindCurrenciesResponse{Currencies: results}, nil
	} else {
		data, err := currencyDao.FindCurrencyValuesByDate(ctx, request.Finit, request.Fend, in.Currency)
		if err != nil {
			return FindCurrenciesResponse{}, err
		}

		results := make(map[string][]model.CurrencyResponse)
		for _, v := range data {
			c := model.CurrencyResponse{Value: v.Value, Date: v.Timestamp.Format("2006-01-02T15:04:05")}
			results[in.Currency] = append(results[in.Currency], c)
		}

		if len(results) == 0 {
			return FindCurrenciesResponse{results}, errs.NotFound
		}

		return FindCurrenciesResponse{results}, nil
	}

}

func requestValidation(in FindCurrenciesRequest) error {
	// Validar la longitud de la moneda
	if len(in.Currency) != 3 {
		return errs.CurrencyInvalidCharacter
	}

	// Validar que los 3 caracteres sean letras y no números
	regex := regexp.MustCompile("^[A-Za-z]{3}$")
	if !regex.MatchString(in.Currency) {
		return errs.CurrencyInvalidName
	}

	return nil
}

func dateValidation(ctx context.Context, in FindCurrenciesRequest) (FindCurrenciesRequest, error) {

	currencyDao := daoCu.NewCurrencyDAO()

	var start time.Time
	var end time.Time

	if in.Finit.String() == "0001-01-01 00:00:00 +0000 UTC" {
		dateInit, err := currencyDao.FindCurrencyValueDate(ctx, in.Currency, "FirstDate")
		if err != nil {
			return FindCurrenciesRequest{}, err
		}
		start = dateInit
	} else {
		start = in.Finit
	}

	if in.Fend.String() == "0001-01-01 00:00:00 +0000 UTC" {
		dateEnd, err := currencyDao.FindCurrencyValueDate(ctx, in.Currency, "LastDate")
		if err != nil {
			return FindCurrenciesRequest{}, err
		}
		end = dateEnd
	} else {
		end = in.Fend
	}

	response := FindCurrenciesRequest{Finit: start, Fend: end}

	return response, nil
}
