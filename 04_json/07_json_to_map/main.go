package main

import (
	"encoding/json"
	"fmt"
)

var JSON_STRING = `
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

func main() {
	var data map[string]any

	err := json.Unmarshal([]byte(JSON_STRING), &data)
	if err != nil {
		panic(err)
	}

	id, _ := data["id"].(float64)

	fmt.Printf("id: %v", id)
}
