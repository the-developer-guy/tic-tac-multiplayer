package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"sync"
)

type AppConfig struct {
	AdminUser     string
	AdminPassword string
	AdminToken    string
}

func loadConfig() {

	ac := AppConfig{
		AdminUser:     os.Getenv("ADMIN_USER"),
		AdminPassword: os.Getenv("ADMIN_PASS"),
		AdminToken:    os.Getenv("ADMIN_TOKEN"),
	}

	if ac.AdminUser == "" {
		panic("Missing admin username from config")
	}
	if ac.AdminPassword == "" {
		panic("Missing admin password from config")
	}
	if ac.AdminToken == "" {
		panic("Missing admin token from config")
	}
}

type TicTacToeServer struct {
	Lobbies     map[string]*Lobby
	lobbiesLock sync.Mutex
}

func NewTicTacToeServer() *TicTacToeServer {
	loadConfig()
	ttts := TicTacToeServer{
		Lobbies: make(map[string]*Lobby),
	}
	return &ttts
}

func (ttts *TicTacToeServer) RegisterApiHandles() {
	http.HandleFunc("GET /playerinfo/{playerId}/", ttts.HandlePlayerInfo)
	http.HandleFunc("GET /ready/{playerId}/", ttts.HandleReadyPlayer)

	http.HandleFunc("GET /getgrid/{lobbyId}/", ttts.HandleGetGameGrid)
	http.HandleFunc("POST /place/{lobbyId}/", ttts.HandlePlaceMark)
	http.HandleFunc("GET /getscores/", ttts.HandleFetchPlayerScores)
	http.HandleFunc("GET /scores/", ttts.HandleScoresView)
}

func (ttts *TicTacToeServer) RegisterAdminHandles() {
	http.HandleFunc("GET /admin/players/", ttts.HandleAdminListPlayers)

	http.HandleFunc("/login/", ttts.HandleLoginView)
	http.HandleFunc("POST /accessc", ttts.HandleAccessControl)
	http.HandleFunc("/adminpage/", ttts.HandleAdminView)
	http.HandleFunc("/fetchdata/", ttts.HandleGetData)
	http.HandleFunc("POST /manualnewplayer/", ttts.HandleNewPlayer)
	http.HandleFunc("POST /regeneratetoken/", ttts.HandleRegenerateToken)
	http.HandleFunc("POST /handleplayeraccess/", ttts.HandleEditPlayerPermissions)
	http.HandleFunc("/playerimport/", ttts.HandleImportPlayerView)
}

func (ttts *TicTacToeServer) AddLobby(lobby *Lobby) {
	ttts.lobbiesLock.Lock()
	// TODO add check if lobby ID exists
	ttts.Lobbies[lobby.LobbyID] = lobby
	ttts.lobbiesLock.Unlock()
}

func (ttts *TicTacToeServer) GetLobby(lobbyId string) (*Lobby, error) {
	return nil, errors.New("Not implemented")
}

func (ttts *TicTacToeServer) Json() ([]byte, error) {
	ttts.lobbiesLock.Lock()
	payload, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	return payload, err
}
