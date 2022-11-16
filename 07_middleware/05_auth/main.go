package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Id   int
	Name string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/foo", FooHandler)

	middlewares := []func(http.Handler) http.Handler{
		LoggingMiddleware,
		SessionMiddleware,
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Write([]byte("Home"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("foo")
	user, _ := r.Context().Value("user").(User)
	fmt.Printf("%+v\n", user)
	w.Write([]byte("Foo"))
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

func SessionMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session")
		if cookie != nil {
			sessionId := cookie.Value
			user := GetUserBySession(sessionId)
			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)
		}
		handler.ServeHTTP(w, r)
	})
}

func GetUserBySession(sessionId string) User {
	if sessionId == "123" {
		return User{Name: "admin"}
	}
	return User{}
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(User)
		adminRequired := strings.HasPrefix(r.URL.Path, "/admin/")
		isAdmin := user.Name != "admin"

		if adminRequired && !isAdmin {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "403 forbidden")
			return
		}

		handler.ServeHTTP(w, r)
	})
}
