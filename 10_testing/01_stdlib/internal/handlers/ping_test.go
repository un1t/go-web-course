package handlers_test

import (
	"example/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handlers.Ping(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	content := w.Body.String()
	expected := "OK"

	if content != expected {
		t.Fatalf(`expected content "%s", got %s`, expected, content)
	}
}
