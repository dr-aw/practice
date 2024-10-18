package main

import (
	"context"
	"fmt"
	"github.com/dr-aw/practice/internal/app"
	"github.com/dr-aw/practice/internal/app/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	_, cancel := context.WithCancel(context.Background())

	app.Run()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connect to the DataBase: %v", err)
	}

	fmt.Println("DB connected successfully:", db)
	//username := "admin"
	//password := "mySecretPassword"
	//if err := database.AuthUser(db, username, password); err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Printf("login succeed: %s", username)
	//}
	if err := database.AddUser(db, "user3", "43214321"); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("user added: %s", "user3")
	}
	<-signalChan
	fmt.Println("Received shutdown signal")
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Shutting down gracefully")
}
