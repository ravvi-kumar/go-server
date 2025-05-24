package main

import (
	"net/http"

	"github.com/ravvi-kumar/go-server/internal/app"
	"github.com/ravvi-kumar/go-server/internal/routes"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
