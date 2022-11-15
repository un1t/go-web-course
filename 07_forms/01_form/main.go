package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/form", FormHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		io.WriteString(w, "405 Method Not Allowed")
		return
	}

	foo := r.FormValue("foo")

	io.WriteString(w, "OK: "+foo)
}
