package auth

import (
	"errors"
	"fmt"
)

type PlayerAuth struct {
	players map[int64]*Player // ID-Player
}

func NewPlayerAuth() *PlayerAuth {
	pa := PlayerAuth{
		players: make(map[int64]*Player),
	}

	return &pa
}

func (pa *PlayerAuth) AddPlayer(id int64, player *Player) error {
	return errors.New("not implemented")
}

func (pa *PlayerAuth) LoadPlayers() error {
	return errors.New("not implemented")
}

func (pa *PlayerAuth) GetPlayer(id int64) (*Player, error) {
	p, ok := pa.players[id]
	if !ok {
		return nil, fmt.Errorf("player ID %d not found", id)
	}

	return p, nil
}

func (pa *PlayerAuth) GetAuthenticatedPlayer(id int64, token string) (*Player, error) {
	p, ok := pa.players[id]
	if !ok {
		return nil, fmt.Errorf("player ID %d not found", id)
	}

	if p.Token != token {
		return nil, fmt.Errorf("invalid token for player ID %d", id)
	}

	return p, nil
}
