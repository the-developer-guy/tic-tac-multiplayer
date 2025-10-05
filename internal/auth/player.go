package auth

import "time"

type PlayerScores struct {
	WinCount  int `json:"winCount"`
	LoseCount int `json:"loseCount"`
	TieCount  int `json:"tieCount"`
}

type Player struct {
	Name           string        `json:"name"`
	Token          string        `json:"token"`
	BanTimestamp   *time.Time    `json:"banTimestamp"`
	DateOfRegister time.Time     `json:"dateOfRegister"`
	Scores         *PlayerScores `json:"scores"`
}

func NewPlayer(name, token string) *Player {
	p := Player{
		Name:           name,
		Token:          token,
		BanTimestamp:   nil,
		DateOfRegister: time.Now(),
		Scores: &PlayerScores{
			WinCount:  0,
			LoseCount: 0,
			TieCount:  0,
		},
	}

	return &p
}
