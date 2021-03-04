package model

import "time"

type Game struct {
	Name          string
	Teams         []Team
	Period        int
	TimeRemaining time.Duration
}

type Team struct {
	Name  string
	Score int
}
