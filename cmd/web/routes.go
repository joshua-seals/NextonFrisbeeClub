package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {
	r := http.NewServeMux()
	// r := chi.NewRouter()
	// Fileserver for assets
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	// General routes go here
	r.HandleFunc("GET /{$}", app.home)
	r.HandleFunc("GET /about", app.about)
	r.HandleFunc("GET /ultimate", app.ultimate)
	return r
}
