package server

import (
	"fmt"
	"net/http"
)

func GetGameGrid(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get arena")
}

func PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")
}
