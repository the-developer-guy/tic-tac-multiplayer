package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (ttts *TicTacToeServer) GetGameGrid(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get arena")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}

	r := ttts.GenerateGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func (ttts *TicTacToeServer) PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}
}

func (ttts *TicTacToeServer) GetLobbyStatus(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting lobby status")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}
}

func (ttts *TicTacToeServer) GenerateGrid() TicTacToeGrid {

	grid := TicTacToeGrid{
		XMarks: []TicTacToeMark{},
		OMarks: []TicTacToeMark{},
	}
	return grid
}

func (ttts *TicTacToeServer) HandleCreateLobby(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminToken := r.Form.Get("atoken")
	if adminToken != "admin" { //TODO: Must be replaced in the future
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	newLobby := CreateLobbyFromRequest(r)

	ttts.lobbiesLock.Lock()
	ttts.Lobbies[newLobby.LobbyID] = newLobby
	ttts.lobbiesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"Lobbyid": newLobby.LobbyID}
	json.NewEncoder(w).Encode(response)

	RegisterLobbyHandlers(newLobby.LobbyID)
}

func HandlePlaceInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
	r.ParseForm()
	token := r.Form.Get("token")
	corX := r.Form.Get("cor_x")
	corY := r.Form.Get("cor_y")
	if token == "" || corX == "" || corY == "" {
		http.Error(w, "Missing Arguments", http.StatusBadRequest)
		return
	}
	fmt.Println("Handling Mark Placement")
}

func HandleGetStatusInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
	fmt.Println("Handling getstatus")
	//lobbyPath is going to come in handy when it comes to actually handling the request.
}

func (ttts *TicTacToeServer) GetActiveLobbies(w http.ResponseWriter, r *http.Request) {
	ttts.lobbiesLock.Lock()
	jsonData, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func RegisterLobbyHandlers(lobbyID string) {
	placePath := fmt.Sprintf("POST /%s/place", lobbyID)
	statusPath := fmt.Sprintf("/%s/getstatus", lobbyID)

	http.HandleFunc(placePath, func(w http.ResponseWriter, r *http.Request) {
		HandlePlaceInLobby(w, r, lobbyID)
	})

	http.HandleFunc(statusPath, func(w http.ResponseWriter, r *http.Request) {
		HandleGetStatusInLobby(w, r, lobbyID)
	})
}

func CreateLobbyFromRequest(req *http.Request) *Lobby {
	req.ParseForm()
	token1 := req.Form.Get("token")
	token2 := req.Form.Get("token2")

	return NewLobby(token1, token2)
}
