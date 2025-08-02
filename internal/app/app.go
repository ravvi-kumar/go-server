package app

import (
	"log"
	"os"
	"time"

	"github.com/ravvi-kumar/go-server/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "[APP] ", log.Ldate|log.Ltime)

	// our stores will go here

	// our api handlers will go here
	workoutHandler := api.NewWorkoutHandler()
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}
	return app, nil
}

func (app *Application) Start() {
	for {
		time.Sleep(time.Second)
		app.Logger.Println("tick")
	}
}
