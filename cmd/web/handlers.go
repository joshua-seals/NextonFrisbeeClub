package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	gameStatus := models.GetStatus()
	date := time.Now().Format("January 2, 2006")
	data := &templateData{
		GameStatus: gameStatus,
		Date:       date,
	}
	app.render(w, http.StatusOK, page, data)
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
	err := r.ParseMultipartForm(10 << 2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := r.FormValue("loginEmail")
	password := r.FormValue("loginPassword")
	log.Printf("Login- Name: %s, Email: %s Pass: %s", email, password)
	// Pull hash from db for comparison
	hashedPassword := []byte("password123")
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Invalid password!")
	} else {
		fmt.Println("Password is correct.")
	}
	// Set session cookie and allow auth or deny

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

	// Here you would typically validate the data and create a new user in your database
	log.Printf("New signup: Name: %s, Email: %s Pass: %s", name, email, password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	// Store in db
	log.Printf(string(hashedPassword))
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

// Temporarily Update the Game Status
func (app *App) gameStatus(w http.ResponseWriter, r *http.Request) {

	statusCode := r.PathValue("status")
	if statusCode == "1" {
		// Cancel the game for 6 hours
		models.SetTemporaryOverride("CANCELED", 12*time.Hour)
		fmt.Fprintln(w, "Game status set to CANCELED for 12 hours.")
	} else {
		// You can optionally support other status codes or ignore them
		fmt.Fprintln(w, "Unrecognized status. No change made.")
	}

}
