package handlers_test

import (
	"example/internal/handlers"
	"example/internal/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPost(t *testing.T) {
	db := tests.SetupDB()
	defer tests.TeardownDB(db)

	r, err := http.NewRequest("GET", "/post/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := handlers.GetPost(db)
	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("status code: got %d want %d", w.Code, http.StatusOK)
	}

	content := w.Body.String()
	t.Log("AAA11", content)
}
