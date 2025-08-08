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

	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	app.Logger.Printf("server is running")

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test ok"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
