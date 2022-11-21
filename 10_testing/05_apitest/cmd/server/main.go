package main

import (
	"example/internal/app"
	"fmt"
)

func main() {
	app := app.NewApp()
	err := app.Config.Load(".env")
	if err != nil {
		panic(err)
	}

	err = app.Setup()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting is running.")

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
