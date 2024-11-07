package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect(){
	// load env file
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the DB URL from the environment variable
    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("DB_URL is not set in the environment")
    }

	d, err := gorm.Open("mysql", dbURL)

	if err != nil{
		panic((err))
	}
	db = d
}

func GetDb() * gorm.DB{
	return db
}