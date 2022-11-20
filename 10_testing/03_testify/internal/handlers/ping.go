package handlers

import (
	"io"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
