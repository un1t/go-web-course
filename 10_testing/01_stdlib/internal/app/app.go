package app

import (
	"errors"
	"example/internal/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
}

func NewApp() App {
	return App{}
}

func (app *App) Configure() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return errors.New("missing DATABASE_URL")
	}

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.HandleFunc("/post", handlers.GetPosts(db)).Methods("GET")
	r.HandleFunc("/post", handlers.CreatePost(db)).Methods("POST")
	r.HandleFunc(`/post/{id:\d+}`, handlers.GetPost(db)).Methods("GET")
	r.HandleFunc(`/post/{id:\d+}`, handlers.UpdatePost(db)).Methods("PUT")
	r.HandleFunc("/post", handlers.DeletePost(db)).Methods("DELETE")
	app.Router = r

	return nil
}

func (app *App) Run() error {
	return http.ListenAndServe(":3000", app.Router)
}
