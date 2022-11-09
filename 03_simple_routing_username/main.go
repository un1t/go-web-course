package main

import (
	"fmt"
	"net/http"
	"strings"
)

type MyHandler struct {
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home"))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/hello/") {
		name := strings.Split(r.URL.Path, "/")[2]
		w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}

func main() {
	http.ListenAndServe(":3000", MyHandler{})
}
