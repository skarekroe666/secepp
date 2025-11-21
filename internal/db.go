package internal

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("secrets.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = db.AutoMigrate(&Secret{})
	if err != nil {
		log.Fatal("failed to migrate to database")
	}

	fmt.Println("migrated successfully")

	DB = db
}
