package models

import (
	"database/sql"
	"time"
)

type Player struct {
	ID          int
	Name        string
	Image       string // url string to retrieve image
	GamesPlayed int
	Offense     int
	Defense     int
	Speed       int
	Endurance   int
	Style       string
	Created     time.Time
	Updated     time.Time
}

type UltimateModel struct {
	DB *sql.DB
}
