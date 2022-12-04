package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id    int    `json:"id" ozzo:"id"`
	Name  string `json:"name" ozzo:"имя"`
	Email string `json:"email" ozzo:"почта"`
	Phone string `json:"phone" ozzo:"телефон"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required,
			validation.Length(2, 50).Error("длинна должна быть от 2 до 50 символов")),
		validation.Field(&u.Email, validation.Required,
			is.Email.Error("неверный адрес почты")),
		validation.Field(&u.Phone, is.E164.Error("неверный номер")),
	)
}

func main() {
	validation.ErrorTag = "ozzo"

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
	if r.Method != http.MethodPost {
		WriteJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "method not allowed",
		})
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	err = user.Validate()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("user %v", user)

	WriteJson(w, http.StatusOK, map[string]any{
		"ok": true,
	})
}
