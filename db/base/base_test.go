package base

import (
	"log"
	"time"

	"testing"

	m "github.com/Alonso-Arias/test-boletia/db/model"
)

func TestGetConnection(t *testing.T) {

	dbc := GetDB()

	result := m.Currency{}

	dbc.Raw("SELECT * FROM CURRENCIES").Scan(&result)

}

func TestGetTime(t *testing.T) {

	loc, _ := time.LoadLocation("Europe/Monaco")
	//set timezone,
	savetrxTime := time.Now().In(loc)

	log.Println("Hora  : ", savetrxTime)

}
