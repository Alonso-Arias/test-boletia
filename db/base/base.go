package base

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := os.Getenv("POSTGRESQL_CONNECTION")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Error connecting to database: error=%v", err)
	}

}

// GetDB gets connection to DB with Gorm
func GetDB() *gorm.DB {
	return db
}
