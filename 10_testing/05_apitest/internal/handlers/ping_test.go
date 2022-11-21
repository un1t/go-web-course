package handlers_test

import (
	"example/internal/handlers"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestPing(t *testing.T) {
	apitest.
		HandlerFunc(handlers.Ping).
		Get("/ping").
		Expect(t).
		Status(http.StatusOK).
		Body("OK").
		End()
}
