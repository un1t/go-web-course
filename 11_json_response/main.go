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
	http.ListenAndServe(":3000", nil)

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
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
