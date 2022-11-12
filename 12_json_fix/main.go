package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", UserHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user1 := User{Name: "Andrey", Id: 555}
	err := WriteJson(w, http.StatusOK, user1)

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}
}
