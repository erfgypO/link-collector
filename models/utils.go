package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var Logger *logrus.Logger

func CreateDbConnection() {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&AccessToken{})
	db.AutoMigrate(&Link{})
	DB = db
}

func SetupLogger(logger *logrus.Logger) {
	Logger = logger
}
