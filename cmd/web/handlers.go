package main

import (
	"net/http"
	"strconv"

	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
)

/*
Need:
 - Login
 - Login with Telegram
Member Features:
 - Create user card
 - Edit owned user card
 - Upload image
 - Rating system for other players
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
func (app *App) SignUp(w http.ResponseWriter, r *http.Request) {

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
