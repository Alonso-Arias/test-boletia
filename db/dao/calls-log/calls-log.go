package dao

import (
	"context"

	"github.com/Alonso-Arias/test-boletia/db/base"
	"github.com/Alonso-Arias/test-boletia/db/model"
	"github.com/Alonso-Arias/test-boletia/log"
)

var loggerf = log.LoggerJSON().WithField("package", "dao")

// CallsLogDAO - CallsLog dao interface
type CallsLogDAO interface {
	Save(ctx context.Context, CallsLog model.CallsLog) error
}

// CallsLogDAOImpl - CallsLog dao implementation
type CallsLogDAOImpl struct {
}

// NewCallsLogDAO - gets an CallsLogDAOImpl instance
func NewCallsLogDAO() *CallsLogDAOImpl {
	return &CallsLogDAOImpl{}
}

// saving data
func (pd *CallsLogDAOImpl) Save(CallsLog model.CallsLog) error {

	log := loggerf.WithField("struct", "CallsLogDAOImpl").WithField("function", "Save")

	db := base.GetDB()

	err := db.Create(&CallsLog)

	if err.Error != nil {
		log.Debugf("%v", err.Error)
		return err.Error
	}

	log.Infof("Save CallsLog Sucessfull\n")

	return nil

}
