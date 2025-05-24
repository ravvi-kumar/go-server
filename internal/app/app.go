package app

import (
	"log"
	"os"
	"time"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "[APP] ", log.Ldate|log.Ltime)
	app := &Application{
		Logger: logger,
	}
	return app, nil
}

func (app *Application) Start() {
	for {
		time.Sleep(time.Second)
		app.Logger.Println("tick")
	}
}
