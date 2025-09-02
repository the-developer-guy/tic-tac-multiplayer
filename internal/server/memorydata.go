package server

import (
	"fmt"
)

type FetchedData struct {
	data map[int]Player
}

// Constructor -> sample data
func NewFetchedData() *FetchedData {
	return &FetchedData{
		data: map[int]Player{
			1: {Name: "Lakatos Tivadar", Token: "123", IsBanned: false, DateOfRegister: "1/1/26"},
			2: {Name: "Zsoric Migmond", Token: "123asd", IsBanned: false, DateOfRegister: "1/1/26"},
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
