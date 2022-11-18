package main

import (
	"example/internal/app"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := app.NewApp()

	err = server.Configure()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting is running.")

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
