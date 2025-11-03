package server

import (
	"net/http"
)

func (gs *GameServer) HandleAdminListPlayers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminToken := r.Form.Get("admintoken")

	if adminToken != gs.settings.AdminToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("ok"))
}
