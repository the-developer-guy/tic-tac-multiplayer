package server

import (
	"fmt"

	"github.com/google/uuid"
)

type FetchedData struct {
	data map[int]Player
}

// Constructor -> sample data
func NewFetchedData() *FetchedData {
	return &FetchedData{
		data: map[int]Player{
			1: {Name: "Lakatos Tivadar", Token: "123", IsBanned: false, DateOfRegister: "1/1/26"},
			2: {Name: "Zsoric Migmond", Token: "123asd", IsBanned: true, DateOfRegister: "1/1/26"},
			3: {Name: "Lakatos Tivadar", Token: "asd123", IsBanned: false, DateOfRegister: "1/1/26"},
		},
	}
}

func (f *FetchedData) GetAllData() map[int]Player {
	return f.data
}

func (f *FetchedData) GetDataByID(id int) (*Player, error) {
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
	new_ID := len(f.data) + 1
	f.data[new_ID] = Player{
		Name:           name,
		Token:          uuid.NewString(),
		IsBanned:       false,
		DateOfRegister: "2/2/26",
	}

}

func (f *FetchedData) RegenerateToken(token string) error {
	id, err := f.GetDataByToken(token)
	if err != nil {
		return fmt.Errorf("Error regenerating token.")
	}
	player := f.data[id]
	player.Token = uuid.NewString()
	f.data[id] = player

	return nil
}
