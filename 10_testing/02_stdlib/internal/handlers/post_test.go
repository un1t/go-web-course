package handlers_test

import (
	"encoding/json"
	"example/internal/app"
	"example/internal/models"
	"example/internal/rest"
	"example/internal/tests"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestGetPost(t *testing.T) {
	var err error

	app := app.NewApp()

	configPath := filepath.Join(tests.GetProjectRoot(), ".env.test")
	err = app.Config.Load(configPath)
	if err != nil {
		t.Fatal(err)
	}

	err = app.Setup()
	if err != nil {
		t.Fatal(err)
	}
	defer app.Teardown()

	db := app.DB

	tests.SetupDB(db)
	defer tests.TeardownDB(db)

	post := models.Post{
		Title: "some title",
		Text:  "some content",
	}
	err = db.Create(&post).Error
	if err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("/post/%d", post.Id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	// handler := handlers.GetPost(db)
	handler := app.Router
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status code: expected %d, got %d", http.StatusOK, w.Code)
	}

	type PostReponse struct {
		rest.Response
		Result models.Post `json:"result"`
	}
	var response PostReponse

	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Result.Title != "some title" {
		t.Fatalf(`title: expected "%s", got "%s"`,
			"some content", response.Result.Title)
	}

	if response.Result.Text != "some content" {
		t.Fatalf(`text: expected "%s", got "%s"`,
			"some content", response.Result.Text)
	}
}
