package base

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {

	//log := loggerf.WithField("func", "init")

	host := os.Getenv("HP_POSTGRES_HOST")
	port := os.Getenv("HP_POSTGRES_PORT")
	user := os.Getenv("HP_POSTGRES_USER")
	password := os.Getenv("HP_POSTGRES_PASSWORD")
	dbname := os.Getenv("HP_POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

}

// GetDB gets connection to DB with Gorm
func GetDB() *gorm.DB {
	return db
}
