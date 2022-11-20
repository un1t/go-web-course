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

	"github.com/stretchr/testify/assert"
)

func TestGetPost(t *testing.T) {
	var err error

	app := app.NewApp()

	configPath := filepath.Join(tests.GetProjectRoot(), ".env.test")
	err = app.Config.Load(configPath)
	assert.Nil(t, err)

	err = app.Setup()
	assert.Nil(t, err)
	defer app.Teardown()

	db := app.DB

	tests.SetupDB(db)
	defer tests.TeardownDB(db)

	post := models.Post{
		Title: "some title",
		Text:  "some content",
	}
	err = db.Create(&post).Error
	assert.Nil(t, err)

	url := fmt.Sprintf("/post/%d", post.Id)

	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	// handler := handlers.GetPost(db)
	handler := app.Router
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type PostReponse struct {
		rest.Response
		Result models.Post `json:"result"`
	}
	var response PostReponse

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "some title", response.Result.Title)
	assert.Equal(t, "some content", response.Result.Text)
}
