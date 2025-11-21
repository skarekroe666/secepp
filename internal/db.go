package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("failed to load .env file")
	}

	url := os.Getenv("DB_URL")

	return url, nil
}

func InitDB() {
	dbUrl, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = db.AutoMigrate(&Secret{})
	if err != nil {
		log.Fatal("failed to migrate to database")
	}

	// fmt.Println("migrated successfully")

	DB = db
}
