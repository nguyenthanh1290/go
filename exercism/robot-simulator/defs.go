package robot

// definitions used in step 1

// Step1Robot is robot.
var Step1Robot struct {
	X, Y int
	Dir
}

// Dir represents directions.
type Dir int

// Directions: North, East, South, West.
const (
	N Dir = iota
	E
	S
	W
)

// String returns formatted direction.
func (d Dir) String() (s string) {
	switch d {
	case N:
		s = "North"
	case E:
		s = "East"
	case S:
		s = "South"
	case W:
		s = "West"
	}
	return
}

// Right turns right.
func Right() {
	Step1Robot.Dir++
	if Step1Robot.Dir > W {
		Step1Robot.Dir = N
	}
}

// Left turns left.
func Left() {
	Step1Robot.Dir--
	if Step1Robot.Dir < N {
		Step1Robot.Dir = W
	}
}

// Advance moves ahead.
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

// additional definitions used in step 2

// Command commands the robot to move. Valid values are 'R', 'L', 'A'.
type Command byte
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

// StartRobot starts robot.
func StartRobot(chan Command, chan Action) {

}

// Room prepares the room.
func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {

}

// additional definition used in step 3

// type Step3Robot struct {
// 	Name string
// 	Step2Robot
// }
