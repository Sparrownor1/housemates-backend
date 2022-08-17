package db

import (
	"fmt"
	"housemates/housemates-backend/core/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Init() {
	var err error

	// get dbinfo from .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", db_user, db_password, db_name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// migrations
	db.AutoMigrate(&models.User{}, &models.Group{}, &models.List{}, &models.ListItem{})
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("db uninitialized")
	}
	return db
}
