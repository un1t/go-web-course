package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo())
	http.HandleFunc("/hello/", hello)

	http.ListenAndServe(":3000", nil)
}

func foo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HanlderFunc"))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, user"))
}
