package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

// ConnectDB uses .env for connection to DB with GORM
func ConnectDB() (*gorm.DB, error) {
	// Read .env
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error load .env file: %v", err)
	//}

	dbHost := os.Getenv("DB_HOST")
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
	pwHash, err := hashPassword(password)
	if err != nil {
		return err
	}

	// Creating new user
	user := User{
		Username:     username,
		PasswordHash: pwHash,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
