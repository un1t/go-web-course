package main

import (
	"encoding/json"
	"fmt"
)

var DATA = `
{
	"id": 55,
	"price": 3000,
	"items": [
		{
			"name": "snowbord",
			"number": 1
		},
		{
			"name": "ball",
			"number": 4
		}
	]
}
`

type Order struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Items []Item `json:"items"`
}

type Item struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func main() {
	var order1 Order

	err := json.Unmarshal([]byte(DATA), &order1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", order1)
}
