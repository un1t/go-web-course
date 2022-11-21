package tests

import (
	"example/internal/app"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func AppSetup(t *testing.T) *app.App {
	app := app.NewApp()

	configPath := filepath.Join(GetProjectRoot(), ".env.test")
	err := app.Config.Load(configPath)
	require.Nil(t, err)

	err = app.Setup()
	require.Nil(t, err)

	MigrateUp(app.DB)

	return &app
}

func AppTeardown(app *app.App) {
	MigrateDown(app.DB)
	app.Teardown()
}
