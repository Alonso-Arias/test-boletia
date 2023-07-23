package dao

import (
	"context"
	"time"

	"github.com/Alonso-Arias/test-boletia/db/base"
	"github.com/Alonso-Arias/test-boletia/db/model"
	"github.com/Alonso-Arias/test-boletia/log"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// CurrencyDAO - Currency dao interface
type CurrencyDAO interface {
	Save(ctx context.Context, Currency model.Currency) error
}

// CurrencyDAOImpl - Currency dao implementation
type CurrencyDAOImpl struct {
}

// NewCurrencyDAO - gets an CurrencyDAOImpl instance
func NewCurrencyDAO() *CurrencyDAOImpl {
	return &CurrencyDAOImpl{}
}

// saving data
func (pd *CurrencyDAOImpl) Save(Currency model.Currency) error {

	log := loggerf.WithField("struct", "CurrencyDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&Currency)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save Currency Sucessfull\n")

	return nil

}

func (pd *CurrencyDAOImpl) FindCurrencyValuesByDate(ctx context.Context, start time.Time, end time.Time, currency string) ([]model.Currency, error) {

	log := loggerf.WithField("struct", "CurrencyDAOImpl").WithField("function", "FindCurrencyValuesByDate")

	db := base.GetDB()

	var currencies []model.Currency

	err := db.Model(currencies).
		Where("code = ? AND timestamp BETWEEN ? AND ?", currency, start.Format("2006-01-02T15:04:05"), end.Format("2006-01-02T15:04:05")).
		Select("value, timestamp").
		Find(&currencies).
		Error
	if err != nil {
		log.WithError(err).Errorf("fails query")
		return currencies, err
	}
	return currencies, nil

}
