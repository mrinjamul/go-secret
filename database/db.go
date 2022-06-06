package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrinjamul/go-secret/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// IsConnected returns the connection status
	IsConnected bool
)

func GetDB() *gorm.DB {
	var db *gorm.DB
	// Get ENV variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	if db == nil {
		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		if dbPort == "" {
			dbPort = "5432"
		}
	}

	// Connect to db
	dest := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dest), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	IsConnected = true
	// Migrate the schema
	db.AutoMigrate(&models.Message{})
	return db
}

// GetTableName returns the table name
func GetTableName() string {
	return "messages"
}