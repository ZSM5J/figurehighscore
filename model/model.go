package model

import (
	"time"
)

type FigureHighScore struct {
	FigureID 	string `json:"figureID"`
	Distance 	int    `json:"distance"`
}

type Player struct {
	Token 		string 	  `json:"token"`
	Registered  time.Time `json:"registred"`
}

type Result struct {
	ResID       string `json:"resID"`
	FigureID 	string `json:"figureID"`
	LapTime 	int    `json:"lapTime"`
	Username  	string `json:"username"`
	Token 		string `json:"token"`
	Trashed     bool   `json:"trashed"`
}


