package auth

import "errors"

type PlayerAuth struct {
	Players map[string]*Player // ID-Player
}

func NewPlayerAuth() *PlayerAuth {
	pa := PlayerAuth{
		Players: make(map[string]*Player),
	}

	return &pa
}

func (pa *PlayerAuth) AddPlayer(id string, player *Player) error {
	return errors.New("not implemented")
}

func (pa *PlayerAuth) LoadPlayers() error {
	return errors.New("not implemented")
}
