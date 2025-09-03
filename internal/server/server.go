package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

type TicTacToeServer struct {
	Lobbies     map[string]*Lobby
	lobbiesLock sync.Mutex
}

func (ttts *TicTacToeServer) GetLobby(lobbyId string) (*Lobby, error) {
	return nil, errors.New("Not implemented")
}

func NewTicTacToeServer() *TicTacToeServer {
	ttts := TicTacToeServer{
		Lobbies: make(map[string]*Lobby),
	}
	return &ttts
}

func (ttts *TicTacToeServer) RegisterHandles() {
	http.HandleFunc("/grid/{lobbyId}/", ttts.GetGameGrid)
	http.HandleFunc("/place/{lobbyId}/", ttts.PlaceMark)
	http.HandleFunc("/status/{lobbyId}/", ttts.GetLobbyStatus)
	http.HandleFunc("/getlobbies/", ttts.GetActiveLobbies)
	http.HandleFunc("POST /createlobby/", ttts.HandleCreateLobby) // TODO automate lobby creation
}

func (ttts *TicTacToeServer) RegisterInterfaceHandlers() {
	http.HandleFunc("/login/", ttts.LoginPage)
	http.HandleFunc("POST /accessc", ttts.accessControl)
	http.HandleFunc("/ainterface/", ttts.ainterFace)
	http.HandleFunc("/fetchdata/", ttts.GetData)
	http.HandleFunc("POST /manualnewplayer/", ttts.NewPlayer)
	http.HandleFunc("POST /regeneratetoken/", ttts.RegenerateToken)
	http.HandleFunc("POST /handleplayeraccess/", ttts.BanOrEnablePlayer)
	http.HandleFunc("POST /removeplayer/", ttts.RemoveUser)
}
func (ttts *TicTacToeServer) AddLobby(lobby *Lobby) {
	ttts.lobbiesLock.Lock()
	// TODO add check if lobby ID exists
	ttts.Lobbies[lobby.LobbyID] = lobby
	ttts.lobbiesLock.Unlock()
}

func (ttts *TicTacToeServer) Json() ([]byte, error) {
	ttts.lobbiesLock.Lock()
	payload, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	return payload, err
}
