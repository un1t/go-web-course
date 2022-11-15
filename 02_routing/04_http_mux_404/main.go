package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler404(w, r)
		return
	}
	w.Write([]byte("Home"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.MatchString(r.URL.Path) {
		handler404(w, r)
		return
	}
	username := strings.Split(r.URL.Path, "/")[2]
	fmt.Fprintf(w, "Hello, %s", username)
}

func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}
