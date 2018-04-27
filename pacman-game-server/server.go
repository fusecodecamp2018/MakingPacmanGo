package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Coordinates struct {
	X int
	Y int
}

var Board = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0},
	{0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
	{0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0},
	{0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0},
	{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var GhostCoordinates = make(map[string]Coordinates)
var Pacman = Coordinates{X: 0, Y: 0}

func startGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> START GAME << ")
	fmt.Println(r.URL.Query())
	GhostCoordinates = make(map[string]Coordinates)
	Pacman = Coordinates{X: 0, Y: 0}

	w.WriteHeader(http.StatusOK)
}

func addGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> ADD GHOST << ")
	fmt.Println(r.URL.Query())

	w.WriteHeader(http.StatusOK)
}

func trackPacman(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> TRACK PACMAN << ")
	fmt.Println(r.URL.Query())
	x, _ := r.URL.Query()["x"]
	y, _ := r.URL.Query()["y"]
	xint, _ := strconv.Atoi(x[0])
	yint, _ := strconv.Atoi(y[0])
	Pacman = Coordinates{X: xint, Y: yint}
	//go GhostMovement.CalculateGhostMoves(Pacman, GhostCoordinates)

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
	GhostCoordinates[ghost[0]] = Coordinates{X: xint, Y: yint}
	w.WriteHeader(http.StatusOK)
}

func moveGhost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> MOVE GHOST << ")
	fmt.Println(r.URL.Query())

	movements := [4]string{"UP", "DOWN", "LEFT", "RIGHT"}

	move := movements[rand.Intn(len(movements))]

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

func getBoard(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/api/startGame", startGame)
	http.HandleFunc("/api/addGhost", addGhost)
	http.HandleFunc("/api/trackGhost", trackGhost)
	http.HandleFunc("/api/trackPacman", trackPacman)
	http.HandleFunc("/api/moveGhost", moveGhost)
	http.HandleFunc("/api/isCaught", isCaught)
	http.HandleFunc("/api/getBoard", getBoard)
	http.Handle("/", http.FileServer(http.Dir("../pacman")))
	http.ListenAndServe(":3000", nil)
}
