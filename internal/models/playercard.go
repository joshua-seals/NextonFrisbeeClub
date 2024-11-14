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

func FetchPlayersFromDB() ([]*Player, error) {
	return players, nil
}

var players = []*Player{
	{
		ID:          1,
		Name:        "Joshua",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     6,
		Defense:     6,
		Endurance:   6,
		Speed:       6,
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
		Style:       "all-around dynamic player, precise handling with excellent team support and top notch endurance",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          3,
		Name:        "Christophe",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     9,
		Endurance:   8,
		Speed:       7,
		Style:       "excellent defensive instincts, unorthodox but effective handler with high level of endurance",
		Created:     time.Now(),
		Updated:     time.Now(),
	},
	{
		ID:          4,
		Name:        "Joe",
		Image:       "N/A",
		GamesPlayed: 10,
		Offense:     7,
		Defense:     6,
		Endurance:   6,
		Speed:       6,
		Style:       "technique at reading and boxing out players on defense with effective trick throws and an extra gear that you won't see coming",
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
		Name:        "bob",
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
