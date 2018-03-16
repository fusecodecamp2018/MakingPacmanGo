package main

import (
	"net/http"
	"fmt"
)

func startGame(w http.ResponseWriter, r *http.Request){
	fmt.Println(">> START GAME << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusOK)
}

func trackPacman(w http.ResponseWriter, r *http.Request){
	fmt.Println(">> TRACK PACMAN << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusOK)
}

func trackGhost(w http.ResponseWriter, r *http.Request){
	fmt.Println(">> TRACK GHOST << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusFound)
}

func moveGhost(w http.ResponseWriter, r *http.Request){
	fmt.Println(">> MOVE GHOST << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusConflict)
}

func isGameOver(w http.ResponseWriter, r *http.Request){
	fmt.Println(">> GAME OVER << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusAccepted)
}

func main() {
	http.HandleFunc("/api/startGame", startGame)
	http.HandleFunc("/api/trackGhost", trackGhost)
	http.HandleFunc("/api/moveGhost", moveGhost)
	http.HandleFunc("/api/trackPacman", trackPacman)
	http.HandleFunc("/api/isGameOver", isGameOver)
	http.Handle("/", http.FileServer(http.Dir("../pacman")))
	http.ListenAndServe(":3000", nil)
}
