package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
)

/*
Need:
 - Login
 - Login with Telegram
Member Features:
 - Create user card
 - Upload image
 - Rating system for other players
*/

func (app *App) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (app *App) ultimate(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/ultimate.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *App) about(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/about.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *App) renderPlayerCards(w http.ResponseWriter, data *templateData) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/player-cards-grid.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) players(w http.ResponseWriter, r *http.Request) {
	// Fetch players from database
	players, err := models.FetchPlayersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p := &templateData{
		Players: players,
	}

	app.renderPlayerCards(w, p)
}

func (app *App) playerForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/player-card.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *App) playerCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "<h1>CardID: %s</h1>", id)
}

func (app *App) regsiterPlayerForm(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, header, err := r.FormFile("Image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file size
	if header.Size > 5*1024*1024 {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Validate file type
	if !strings.HasPrefix(header.Header.Get("Content-Type"), "image/") {
		http.Error(w, "File is not an image", http.StatusBadRequest)
		return
	}

	// Create a new filename
	filename := filepath.Join("uploads", generateSecureFilename())

	// Create the file
	dst, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sanitize text inputs
	// https://pkg.go.dev/github.com/microcosm-cc/bluemonday#section-readme
	// p := bluemonday.UGCPolicy()
	// name := p.Sanitize(r.FormValue("Name"))
	// style := p.Sanitize(r.FormValue("Style"))

	// Process other form fields...

	// Redirect or respond as needed
}

func generateSecureFilename() string {
	return ""
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {}
