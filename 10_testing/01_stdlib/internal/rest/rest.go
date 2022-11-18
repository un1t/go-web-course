package rest

import (
	"encoding/json"
	"io"
)

type Response struct {
	Ok     bool   `json:"ok"`
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func WriteJSON(w io.Writer, status int, v any) {
	bytes, _ := json.Marshal(v)
	w.Write(bytes)
}

func WriteError(w io.Writer, status int, err error) {
	WriteJSON(w, status, Response{
		Ok:    false,
		Error: err.Error(),
	})
}
