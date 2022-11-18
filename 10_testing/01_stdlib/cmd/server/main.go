package main

import (
	"example/internal/app"
	"fmt"
)

func main() {
	server := app.NewApp()

	err := server.Configure()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting is running.")

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
