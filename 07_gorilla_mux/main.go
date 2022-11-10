package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/hello/{username}", hello)
	r.HandleFunc(`/product/{id:\d+}`, product)
	r.HandleFunc(`/form`, form).Methods("POST", "PUT")
	r.NotFoundHandler = http.HandlerFunc(handler404)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Fprintf(w, "Hello %s!", username)
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Product ID %s", id)
}

func form(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Form")
}

func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 Page Not Found")
}
