package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/getme", GetMeHandler)

	middlewares := []func(http.Handler) http.Handler{
		LoggingMiddleware,
		AuthMiddleware,
	}

	handler := http.Handler(mux)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id   int
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, map[string]any{
		"ok": true,
	})
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("user").(User)

	if user.Id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		WriteJSON(w, map[string]any{
			"ok":    false,
			"error": "unauthorized",
		})
		return
	}

	WriteJSON(w, map[string]any{
		"ok":   true,
		"user": user,
	})
}

func WriteJSON(w io.Writer, v any) {
	bytes, _ := json.Marshal(v)
	w.Write(bytes)
}

type MyResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *MyResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w2 := &MyResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		handler.ServeHTTP(w2, r)
		log.Printf("%s [%d]\n", r.RequestURI, w2.StatusCode)
	})
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session")
		if cookie != nil {
			sessionId := cookie.Value
			user, _ := GetUserBySessionId(sessionId)

			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)
			r = r.WithContext(ctx)
		}

		handler.ServeHTTP(w, r)
	})
}

func GetUserBySessionId(sessionId string) (User, error) {
	if sessionId == "123" {
		return User{Id: 1, Name: "admin"}, nil
	}
	return User{}, errors.New("session not found")
}
