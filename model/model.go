package model

import (
	"time"
)

type FigureHighScore struct {
	FigureID 	string
	Distance 	int
}

type Player struct {
	Token 		string
	Registered  time.Time
}

type Result struct {
	ResID       string
	FigureID 	string
	LapTime 	int
	Username  	string
	Token 		string
}


