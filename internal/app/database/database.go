package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	Password_Hash string `gorm:"not null"`
}

// ConnectDB uses .env for connection to DB with GORM
func ConnectDB() (*gorm.DB, error) {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file: %v", err)
	}

	dbHost := "192.168.0.154"
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	return db, nil
}

func AddUser(db *gorm.DB, username, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Creating new user
	user := User{
		Username:      username,
		Password_Hash: string(passwordHash),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
