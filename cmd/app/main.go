package main

import (
	"fmt"
	"github.com/dr-aw/practice/internal/app"
	"github.com/dr-aw/practice/internal/app/database"
	"log"
)

func main() {
	app.Run()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect to the DataBase: %v", err)
	}

	fmt.Println("DB connected successfully:", db)
	username := "admin"
	password := "mySecretPassword"
	if err := database.AuthUser(db, username, password); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("login succeed: %s", username)
	}
}
