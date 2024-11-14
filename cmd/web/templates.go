package main

import (
	"github.com/joshua-seals/NextonFrisbeeClub/internal/models"
)

// Template cache

// Add all templateData structs to this struct
// This way it is all accessible.
type templateData struct {
	PlayerCard *models.Player
	Players    []*models.Player
}

// type Templates struct {
// 	templates *template.Template
// }
