package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Farmaan-Malik/Go-Templating/pkg/config"
	"github.com/Farmaan-Malik/Go-Templating/pkg/handlers"
	"github.com/Farmaan-Malik/Go-Templating/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.UseCache = false
	app.TemplateCache = tc
	repo:= handlers.NewRepository(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Server started on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
