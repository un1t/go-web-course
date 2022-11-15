package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Air"))
}

func main() {
	err := http.ListenAndServe(":3000", http.HandlerFunc(index))
	if err != nil {
		panic(err)
	}
}
