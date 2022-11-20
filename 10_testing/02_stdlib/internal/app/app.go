package app

import (
	"errors"
	"example/internal/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Config Config
	Router *mux.Router
	DB     *gorm.DB
}

type Config struct {
	DatabaseUrl string
}

func (config *Config) Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return errors.New("Error loading " + filename)
	}

	config.DatabaseUrl = os.Getenv("DATABASE_URL")
	if config.DatabaseUrl == "" {
		return errors.New("missing DATABASE_URL")
	}

	return nil
}

func NewApp() App {
	return App{}
}

func (app *App) Setup() error {
	db, err := gorm.Open(postgres.Open(app.Config.DatabaseUrl), &gorm.Config{})
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
	app.DB = db

	return nil
}

func (app *App) Teardown() error {
	sqlDB, err := app.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (app *App) Run() error {
	return http.ListenAndServe(":3000", app.Router)
}
