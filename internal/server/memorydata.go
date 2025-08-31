package server

import (
	"fmt"
	"time"
)

type FetchedData struct {
	data map[int]Player
}

// Constructor -> sample data
func NewFetchedData() *FetchedData {
	return &FetchedData{
		data: map[int]Player{
			1: {Name: "Lakatos Tivadar", Token: "123", isBanned: false, dateofRegister: time.Now()},
			2: {Name: "Zsoric Migmond", Token: "123asd", isBanned: false, dateofRegister: time.Now()},
			3: {Name: "Lakatos Tivadar", Token: "asd123", isBanned: false, dateofRegister: time.Now()},
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
