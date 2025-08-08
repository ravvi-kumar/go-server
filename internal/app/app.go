package app

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/ravvi-kumar/go-server/internal/api"
	"github.com/ravvi-kumar/go-server/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {

	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "[APP] ", log.Ldate|log.Ltime)

	// our stores will go here

	// our api handlers will go here
	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	return app, nil
}

func (app *Application) Start() {
	for {
		time.Sleep(time.Second)
		app.Logger.Println("tick")
	}
}
