package GhostMovement

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

func CalculateGhostMoves(pacman Coordinates, ghosts map[string]Coordinates) []string {
	for ghostname, ghostcoord := range ghosts {
		switch ghostname {
		case "blinky":
			go Blinky(pacman, ghostcoord)
		case "clyde":
			go Clyde(pacman, ghostcoord)
		case "funky":
			go Funky(pacman, ghostcoord)
		case "inky":
			go Inky(pacman, ghostcoord)
		case "pinky":
			go Pinky(pacman, ghostcoord)
		}
	}

}

func Blinky(pacman Coordinates, ghost Coordinates) string {
	//always check coordinate below if 1
	//if below 1 move down
	//if left 1 move left
	//if top 1 move up
	//if right 1 move right
	case

}
func Clyde(pacman Coordinates, ghost Coordinates){

}
func Funky(pacman Coordinates, ghost Coordinates){

}
func Inky(pacman Coordinates, ghost Coordinates){

}
func Pinky(pacman Coordinates, ghost Coordinates){

}