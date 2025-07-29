package server

import (
	"encoding/json"
	"net/http"
)

type Player struct {
	Token string //`json: "token" `
	Mark  string //`json: "mark" `
}

type Lobby struct {
	Players []Player      `json:"players"`
	LobbyID string        `json:"lobbyid"`
	Grid    TicTacToeGrid `json:"gamegrid"`
}

func CreateLobby(w http.ResponseWriter, req *http.Request) Lobby {
	lobby := Lobby{
		Players: []Player{
			{Token: "token1", Mark: "X"},
			{Token: "Token2", Mark: "O"},
		},
		LobbyID: "lobbyID123",
		Grid:    writeGridJson(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lobby)

	return lobby
}
