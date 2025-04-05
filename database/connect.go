package database

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/revacholiere-moralist/blogbackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Connected successfully")
	}
	log.Println("this")
	DB = database
	database.AutoMigrate(
		&models.User{},
	)
}
