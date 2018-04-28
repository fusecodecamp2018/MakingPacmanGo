package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/cahcommercial/MakingPacmanGo/pacman-game-server/common"
	"github.com/cahcommercial/MakingPacmanGo/pacman-game-server/GhostMovement"
)

var GhostChannels = make(map[string]chan string)
var GhostCoordinates = make(map[string]*common.Coordinates)
var Pacman = common.Coordinates{X: 0, Y: 0}

func startGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> START GAME << ")
	fmt.Println(r.URL.Query())
	GhostCoordinates = make(map[string]*common.Coordinates)
	GhostChannels = make(map[string]chan string)
	Pacman = common.Coordinates{X: 0, Y: 0}

	w.WriteHeader(http.StatusOK)
}

func addGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> ADD GHOST << ")
	fmt.Println(r.URL.Query())

	ghost, _ := r.URL.Query()["ghost"]
	GhostChannels[ghost[0]] = make(chan string, 2)
	coordinates := common.Coordinates{X: 1000, Y: 1000}
	GhostCoordinates[ghost[0]] = &coordinates

	go GhostMovement.CalculateGhostMoves(&Pacman, ghost[0], &coordinates, GhostChannels[ghost[0]])

	w.WriteHeader(http.StatusOK)
}

func trackPacman(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> TRACK PACMAN << ")
	fmt.Println(r.URL.Query())
	x, _ := r.URL.Query()["x"]
	y, _ := r.URL.Query()["y"]
	xint, _ := strconv.Atoi(x[0])
	yint, _ := strconv.Atoi(y[0])
	Pacman = common.Coordinates{X: xint, Y: yint}

	fmt.Println("Pacman's coordinates are: " + x[0] + "," + y[0])

	w.WriteHeader(http.StatusOK)
}

func trackGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> TRACK GHOST << ")
	fmt.Println(r.URL.Query())
	x, _ := r.URL.Query()["x"]
	y, _ := r.URL.Query()["y"]
	xint, _ := strconv.Atoi(x[0])
	yint, _ := strconv.Atoi(y[0])
	ghost, _ := r.URL.Query()["ghost"]
	fmt.Println("Ghost " + ghost[0] + " coordinates are: " + x[0] + "," + y[0])
	coordinates := GhostCoordinates[ghost[0]]
	coordinates.X = xint
	coordinates.Y = yint
	w.WriteHeader(http.StatusOK)
}

func moveGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> MOVE GHOST << ")
	fmt.Println(r.URL.Query())
	ghost, _ := r.URL.Query()["ghost"]

	fmt.Println("BEFORE MOVE GHOST")
	move := <- GhostChannels[ghost[0]]
	fmt.Println("AFTER MOVE GHOST :" + move)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(move))
}

func isCaught(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(">> GAME OVER << ")
	//fmt.Println(r.URL.Query())
	for _, coordinates := range GhostCoordinates {
		if coordinates.X == Pacman.X && coordinates.Y == Pacman.Y {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("true"))
		}
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/startGame", startGame)
	http.HandleFunc("/api/addGhost", addGhost)
	http.HandleFunc("/api/trackGhost", trackGhost)
	http.HandleFunc("/api/trackPacman", trackPacman)
	http.HandleFunc("/api/moveGhost", moveGhost)
	http.HandleFunc("/api/isCaught", isCaught)
	http.Handle("/", http.FileServer(http.Dir("../pacman")))
	http.ListenAndServe(":3000", nil)
}
