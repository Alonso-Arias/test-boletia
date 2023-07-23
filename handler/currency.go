package handler

import (
	"context"
	"net/http"
	"time"

	errs "github.com/Alonso-Arias/test-boletia/errors"
	"github.com/Alonso-Arias/test-boletia/services/currency"
	"github.com/apex/log"
	"github.com/labstack/echo/v4"
)

func CurrenciesGet(c echo.Context) error {
	format := "2006-01-02T15:04:05"

	var finit time.Time
	var fend time.Time

	if c.QueryParam("finit") != "" {
		dateTime, err := time.Parse(format, c.QueryParam("finit"))
		if err != nil {
			log.WithError(err).Error("Binding error")
			return c.JSON(http.StatusBadRequest, err)
		}
		finit = dateTime
	} else {
		finit = time.Time{} // 0001-01-01 00:00:00 +0000 UTC
	}

	// Mismo procedimiento para c.QueryParam("fend")

	if c.QueryParam("fend") != "" {
		dateTime, err := time.Parse(format, c.QueryParam("fend"))
		if err != nil {
			log.WithError(err).Error("Binding error")
			return c.JSON(http.StatusBadRequest, err)
		}
		fend = dateTime
	} else {
		fend = time.Time{} // 0001-01-01 00:00:00 +0000 UTC
	}

	req := currency.FindCurrenciesRequest{
		Currency: c.Param("currency"),
		Finit:    finit,
		Fend:     fend,
	}

	res, err := currency.CurrencyService{}.FindCurrencies(context.TODO(), req)
	if ce, ok := err.(errs.CustomError); ok {
		return c.JSON(ce.Code, err)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
