package handlers_test

import (
	"example/internal/models"
	"example/internal/tests"
	"fmt"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
)

func TestGetPost(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	post := models.Post{
		Title: "some title",
		Text:  "some content",
	}
	err := db.Create(&post).Error
	assert.Nil(t, err)

	type PostResponse struct {
		Result models.Post `json:"result"`
	}

	url := fmt.Sprintf("/post/%d", post.Id)

	apitest.
		Handler(app.Router).
		Get(url).
		Expect(t).
		Status(http.StatusOK).
		// Body(`{
		// 	"ok": true,
		// 	"result": {
		// 		"id":1,
		// 		"text":"some content",
		// 		"title":"some title"
		// 	}
		// }`).
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result.title", "some title").
				Equal("result.text", "some content").
				End(),
		).
		// Assert(func(resp *http.Response, req *http.Request) error {
		// 	var postResponse PostResponse
		// 	err := json.NewDecoder(resp.Body).Decode(&postResponse)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	assert.Equal(t, "some title", postResponse.Result.Title)
		// 	assert.Equal(t, "some content", postResponse.Result.Text)
		// 	assert.True(t, postResponse.Result.CreatedAt.Before(time.Now()))

		// 	return nil
		// }).
		End()
}
