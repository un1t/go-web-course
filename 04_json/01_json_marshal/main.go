package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	user1 := User{Name: "Ivan", Id: 555}
	bytes, _ := json.Marshal(user1)
	fmt.Println(string(bytes))
}
