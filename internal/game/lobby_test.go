package game

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestLobbyJson(t *testing.T) {
	l := NewLobby("aToken", "bToken", 1, 2, time.Now())
	jsonData, err := json.Marshal(l)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Lobby JSON: %s", string(jsonData))
}
