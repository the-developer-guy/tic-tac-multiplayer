package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLobbyJson(t *testing.T) {
	l := NewLobby("aToken", "bToken")
	jsonData, err := json.Marshal(l)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Lobby JSON: %s", string(jsonData))
}
