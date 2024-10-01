package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *App) Routes() *chi.Mux {
	r := chi.NewRouter()
	// Fileserver for assets
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileserver))
	// General routes go here
	r.Get("/", app.home)
	r.Get("/about", app.about)
	r.Get("/ultimate", app.ultimate)
	return r
}
