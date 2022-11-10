package main

import (
	"net/http"
)

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home"))
		return
	}

	if r.URL.Path == "/hello" {
		w.Write([]byte("Hello, user"))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}

func main() {
	err := http.ListenAndServe(":3000", MyHandler{})
	if err != nil {
		panic(err)
	}
}
