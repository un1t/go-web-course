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

	url := fmt.Sprintf("/post/%d", post.Id)

	apitest.
		Handler(app.Router).
		Get(url).
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result.title", "some title").
				Equal("result.text", "some content").
				End(),
		).
		End()
}
