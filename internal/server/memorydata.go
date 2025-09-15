package server

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FetchedData struct {
	data map[int]Player
}

// Constructor -> sample data
func NewFetchedData() *FetchedData {
	return &FetchedData{
		data: map[int]Player{
			1: {Name: "Lakatos Tivadar", Token: "123", IsBanned: nil, DateOfRegister: "1/1/26"},
			2: {Name: "Zsoric Migmond", Token: "123asd", IsBanned: nil, DateOfRegister: "1/1/26"},
			3: {Name: "Lakatos Tivadar", Token: "asd123", IsBanned: nil, DateOfRegister: "1/1/26"},
		},
	}
}

func (f *FetchedData) GetAllData() map[int]Player {
	return f.data
}

func (f *FetchedData) GetPlayerByID(id int) (*Player, error) {
	if player, ok := f.data[id]; ok {
		return &player, nil
	}
	return nil, fmt.Errorf("Player is not found.")
}

func (f *FetchedData) GetDataByToken(token string) (int, error) {
	for id, player := range f.data {
		if player.Token == token {
			return id, nil
		}
	}
	return 0, fmt.Errorf("No player found with this Token.")
}

func (f *FetchedData) NewPlayer(name string) {
	newID := 0
	for id := range f.data {
		if id > newID {
			newID = id
		}
	}
	newID++

	f.data[newID] = Player{
		Name:           name,
		Token:          uuid.NewString(),
		IsBanned:       nil,
		DateOfRegister: time.Now().UTC().Format(time.RFC3339),
	}
}
func (f *FetchedData) RegenerateToken(token string) (string, error) {
	id, err := f.GetDataByToken(token)
	if err != nil {
		return "", fmt.Errorf("Error regenerating token.")
	}
	player := f.data[id]
	new_token := uuid.NewString()
	player.Token = new_token
	f.data[id] = player

	return new_token, nil
}

func (f *FetchedData) ValidatePlayerAccess(token string) error {
	id, err := f.GetDataByToken(token)
	if err != nil {
		return fmt.Errorf("Error interacting with Player")
	}
	player := f.data[id]
	if player.IsBanned != nil {
		player.IsBanned = nil
	} else {
		now := time.Now()
		player.IsBanned = &now
	}

	f.data[id] = player
	return nil
}
