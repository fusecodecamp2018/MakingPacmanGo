package main

import (
"net/http"
)

func startGame(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/startGame", startGame)
	http.Handle("/", http.FileServer(http.Dir("../pacman")))
	http.ListenAndServe(":3000", nil)
}
