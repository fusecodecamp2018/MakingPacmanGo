package GhostMovement

import (
	"math/rand"
	"github.com/cahcommercial/MakingPacmanGo/pacman-game-server/common"
	"fmt"
)

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

func CalculateGhostMoves(pacman *common.Coordinates, ghostname string, ghostcoord *common.Coordinates, ghostChannel chan string) {
	fmt.Println(">>> CALCULATE GHOST <<<")
	for {
		//TODO fix me cannot get ghosts to have the value from server.go
		fmt.Printf(">>> INFINITE LOOP  <<< %d, %d\n", ghostcoord.X, ghostcoord.Y)
			fmt.Println(">>> DOING IT <<" + ghostname)

			switch ghostname {
			case "blinky":
				fmt.Println(">>> BLINKY <<")
				Blinky(*pacman, *ghostcoord, ghostChannel)
			case "clyde":
				Clyde(*pacman, *ghostcoord, ghostChannel)
			case "funky":
				Funky(*pacman, *ghostcoord, ghostChannel)
			case "inky":
				Inky(*pacman, *ghostcoord, ghostChannel)
			case "pinky":
				Pinky(*pacman, *ghostcoord, ghostChannel)
			default :
				fmt.Println(">>> NOTHING <<")
			}
	}
}

func Blinky(pacman common.Coordinates, ghost common.Coordinates, ghostChannel chan string) {
	//always check coordinate below if 1
	//if below 1 move down
	//if left 1 move left
	//if top 1 move up
	//if right 1 move right
	ghostChannel <- random()
}
func Clyde(pacman common.Coordinates, ghost common.Coordinates, ghostChannel chan string){
	ghostChannel <- random()
}
func Funky(pacman common.Coordinates, ghost common.Coordinates, ghostChannel chan string){
	ghostChannel <- random()
}
func Inky(pacman common.Coordinates, ghost common.Coordinates, ghostChannel chan string){
	ghostChannel <- random()
}
func Pinky(pacman common.Coordinates, ghost common.Coordinates, ghostChannel chan string){
	ghostChannel <- random()
}

func random() string{
	fmt.Println(">>> MOVING <<")
	movements := [4]string{"UP", "DOWN", "LEFT", "RIGHT"}
	return movements[rand.Intn(len(movements))]
}