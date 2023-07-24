package dao

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Alonso-Arias/test-boletia/db/base"
	"github.com/Alonso-Arias/test-boletia/db/model"
	"github.com/Alonso-Arias/test-boletia/log"
	"gorm.io/gorm"
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
		Where("timestamp BETWEEN ? AND ?", start.Format("2006-01-02T15:04:05"), end.Format("2006-01-02T15:04:05")).
		Scopes(CurrencyFilter(strings.ToUpper(currency))).
		Select("currency, value, timestamp").
		Find(&currencies).
		Error
	if err != nil {
		log.WithError(err).Errorf("fails query")
		return currencies, err
	}
	return currencies, nil

}

func (pd *CurrencyDAOImpl) FindCurrencyValueDate(ctx context.Context, currency string, dateType string) (time.Time, error) {
	log := loggerf.WithField("struct", "CurrencyDAOImpl").WithField("function", "FindCurrencyValueDate")

	db := base.GetDB()

	var date time.Time
	var order string

	switch dateType {
	case "FirstDate":
		order = "ASC"
	case "LastDate":
		order = "DESC"
	default:
		return time.Time{}, fmt.Errorf("invalid DateType")
	}

	err := db.Model(&model.Currency{}).
		Scopes(CurrencyFilter(strings.ToUpper(currency))).
		Order("timestamp "+order).
		Limit(1).
		Pluck("timestamp", &date).
		Error
	if err != nil {
		log.WithError(err).Errorf("failed to query for date")
		return time.Time{}, err
	}

	return date, nil
}

func CurrencyFilter(currency string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if currency == "ALL" {
			return db.Where("")
		} else {
			return db.Where("currency = ?", currency)
		}
	}
}

func (pd *CurrencyDAOImpl) GetAllCurrenciesWithLatestValue(ctx context.Context) ([]model.Currency, error) {
	log := loggerf.WithField("struct", "CurrencyDAOImpl").WithField("function", "GetAllCurrenciesWithLatestValue")

	db := base.GetDB()

	var currencies []model.Currency

	// Subconsulta para obtener el valor más reciente para cada divisa
	subquery := db.Table("currencies").
		Select("currency, MAX(timestamp) AS timestamp").
		Group("currency")

	// Unir la subconsulta con la tabla principal para obtener los valores correspondientes
	err := db.Table("currencies").
		Joins("JOIN (?) AS latest ON currencies.currency = latest.currency AND currencies.timestamp = latest.timestamp", subquery).
		Select("currencies.*").
		Find(&currencies).
		Error
	if err != nil {
		log.WithError(err).Errorf("failed to query currencies")
		return nil, err
	}

	return currencies, nil
}
