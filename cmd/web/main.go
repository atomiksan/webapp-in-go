package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atomiksan/webapp-in-go/pkg/config"
	"github.com/atomiksan/webapp-in-go/pkg/handlers"
	"github.com/atomiksan/webapp-in-go/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting app on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
