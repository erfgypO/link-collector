package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

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
