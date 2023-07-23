package currency

import (
	"context"
	"regexp"
	"time"

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
	Currencies model.CurrencyData
}

func (cs CurrencyService) FindCurrencies(ctx context.Context, in FindCurrenciesRequest) (FindCurrenciesResponse, error) {
	// log := loggerf.WithField("service", "CurrencyService").WithField("func", "FindCurrencies")

	// if in.Sku == "" {
	// 	return GetProductResponse{}, errs.BadRequest
	// }

	// productDao := dao.NewProductDAO()
	// productImageDao := dao.NewProductImageDAO()

	// v, err := productDao.Get(ctx, in.Sku)
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	log.WithError(err).Error("problems with getting products")
	// 	return GetProductResponse{}, err
	// } else if err == gorm.ErrRecordNotFound {
	// 	return GetProductResponse{}, errs.ProductsNotFound
	// }

	// pi, err := productImageDao.FindAll(ctx, v.Sku)
	// if err != nil {
	// 	return GetProductResponse{}, err
	// }
	// var productsImages []string
	// for _, item := range pi {
	// 	productsImages = append(productsImages, item.Url)
	// }
	// product := model.Product{
	// 	Sku:              v.Sku,
	// 	Name:             v.Name,
	// 	Brand:            v.Brand,
	// 	Size:             v.Size,
	// 	Price:            v.Price,
	// 	PrincipalImage:   v.PrincipalImage,
	// 	AdditionalImages: productsImages,
	// }

	// return GetProductResponse{Product: product}, nil
	return FindCurrenciesResponse{}, nil
}

func RequestValidation(in FindCurrenciesRequest) error {
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
