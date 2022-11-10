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
	ID    int64 `json:"id"`
	Items []struct {
		Name   string `json:"name"`
		Number int64  `json:"number"`
	} `json:"items"`
	Price int64 `json:"price"`
}

func main() {
	var order1 Order
	err := json.Unmarshal([]byte(DATA), &order1)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(order1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
