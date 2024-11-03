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

func FetchPlayersFromDB() ([]Player, error) {
	return players, nil
}

var players = []Player{
	{
		ID:          1,
		Name:        "Joshua",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     6,
		Defense:     5,
		Endurance:   6,
		Speed:       5,
		Style:       "",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          2,
		Name:        "Ben",
		Image:       "N/A",
		GamesPlayed: 11,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          3,
		Name:        "Bob",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          4,
		Name:        "jon",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          5,
		Name:        "zac",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          6,
		Name:        "zach",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          7,
		Name:        "christophe",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          8,
		Name:        "steven",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     5,
		Endurance:   8,
		Speed:       7,
		Style:       "all-around, precise handling with excellent team support",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
}
