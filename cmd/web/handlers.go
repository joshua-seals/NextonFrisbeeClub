package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
)

/*
Need:
 - Signup (associate name with player card)
 - Login
 - Login with Telegram https://dev.to/shaggyrec/telegram-oauth-authorization-for-your-site-3f4l
Member Features:
 - Edit owned user card
 - Rating system for other players [array- ribbon rack of badges]
 - Badges
*/

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	page := "./ui/html/pages/home.html"
	app.render(w, http.StatusOK, page, nil)
}

func (app *App) ultimate(w http.ResponseWriter, r *http.Request) {
	page := "./ui/html/pages/ultimate.html"
	app.render(w, http.StatusOK, page, nil)
}

func (app *App) about(w http.ResponseWriter, r *http.Request) {
	page := "./ui/html/pages/about.html"
	app.render(w, http.StatusOK, page, nil)
}

func (app *App) players(w http.ResponseWriter, r *http.Request) {
	// Fetch players from database
	players, err := models.FetchPlayersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// stash players in templateData
	data := &templateData{
		Players: players,
	}
	// Render the page with the data
	page := "./ui/html/pages/player-cards-grid.html"
	app.render(w, http.StatusOK, page, data)
}

func (app *App) playerForm(w http.ResponseWriter, r *http.Request) {
	page := "./ui/html/pages/player-form.html"
	// First check to see if user is in Database already

	// Give SignUp
	app.render(w, http.StatusOK, page, nil)
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {

}

// SignUp should immediately take the user to Create New Player
func (app *App) signUp(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get form values
	name := r.FormValue("signupName")
	email := r.FormValue("signupEmail")
	password := r.FormValue("signupPassword")

	// log.Printf("Form: %+v", r.PostForm)
	// log.Printf("MultipartForm: %+v", r.MultipartForm)
	// Here you would typically validate the data and create a new user in your database
	log.Printf("New signup: Name: %s, Email: %s Pass: %s", name, email, password)
	// upon successful signup, direct new user to create player-card
	http.Redirect(w, r, "/players/new", http.StatusSeeOther)
}

func (app *App) playerCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pid, _ := strconv.Atoi(id)
	player := models.FetchPlayerByID(pid)
	if player == nil {
		return
	}
	// hacky for now
	p := []*models.Player{}
	p = append(p, player)
	// stash players in templateData
	data := &templateData{
		Players: p,
	}
	// Render the page with the data
	page := "./ui/html/pages/player-cards-grid.html"
	app.render(w, http.StatusOK, page, data)
}
