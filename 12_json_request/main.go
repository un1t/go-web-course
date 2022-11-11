package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", GetUserHandler).Methods("GET")
	r.HandleFunc("/user", CreateUserHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result := map[string]any{
			"ok":    false,
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	fmt.Printf("user %v\n", user)
	w.Write([]byte("Created!"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user1 := User{Name: "Ivan", Id: 555}
	bytes, err := json.Marshal(user1)

	// err = errors.New("some error")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result := map[string]any{
			"ok":    false,
			"error": err.Error(),
		}
		json.NewEncoder(w).Encode(result)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
