package Lesson_3

import (
	"net/http"
	"fmt"
)

func startGame(w http.ResponseWriter, r *http.Request) {

}

func trackPacman(w http.ResponseWriter, r *http.Request) {

}

func addGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> ADD GHOST << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusOK)
}

func trackGhost(w http.ResponseWriter, r *http.Request) {

}

func moveGhost(w http.ResponseWriter, r *http.Request) {

}

func isGameOver(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/api/startGame", startGame)
	http.HandleFunc("/api/addGhost", addGhost)
	http.HandleFunc("/api/trackGhost", trackGhost)
	http.HandleFunc("/api/moveGhost", moveGhost)
	http.HandleFunc("/api/trackPacman", trackPacman)
	http.HandleFunc("/api/isGameOver", isGameOver)
	http.Handle("/", http.FileServer(http.Dir("../pacman")))
	http.ListenAndServe(":3000", nil)
}
