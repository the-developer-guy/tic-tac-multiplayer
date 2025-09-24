package game

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
