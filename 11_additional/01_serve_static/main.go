package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	staticHandler := http.StripPrefix("/static/", fs)

	http.HandleFunc("/ping", PingHandler)
	http.Handle("/static/", staticHandler)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
