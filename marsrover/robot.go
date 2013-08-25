package marsrover

import (
	"fmt"
)

type Direction string

const (
	N Direction = "N"
	E Direction = "E"
	W Direction = "W"
	S Direction = "S"
)

type Robot struct {
	p         *Planet
	x         int
	y         int
	direction Direction
}

func NewRobot(p *Planet, x int, y int, direction Direction) *Robot {
	return &Robot{p, x, y, direction}
}

func (r *Robot) ExecuteCommand(command string) {
	switch command {
	case "M":
		r.move()
	case "L":
		r.turnLeft()
	case "R":
		r.turnRight()
	default:
	}
}

func (r *Robot) move() {
	targetX, targetY := r.x, r.y

	switch r.direction {
	case N:
		targetY++
	case E:
		targetX++
	case W:
		targetX--
	case S:
		targetY--
	}
	if r.p.IsInBounds(targetX, targetY) {
		r.x, r.y = targetX, targetY
	}
}

func (r *Robot) turnLeft() {
	switch r.direction {
	case N:
		r.direction = W
	case E:
		r.direction = N
	case W:
		r.direction = S
	case S:
		r.direction = E
	}
}

func (r *Robot) turnRight() {
	switch r.direction {
	case N:
		r.direction = E
	case E:
		r.direction = S
	case W:
		r.direction = N
	case S:
		r.direction = W
	}
}

func (r *Robot) Location() string {
	return fmt.Sprintf("%v %v %v", r.x, r.y, r.direction)
}
