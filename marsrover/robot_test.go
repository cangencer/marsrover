package marsrover

import (
	"testing"
)

func assertRobot(t *testing.T,
	initial Robot,
	expected Robot,
	action func(*Robot)) {

	action(&initial)

	if initial.Location() != expected.Location() {
		t.Errorf("Expected state %v, actual state %v", expected.Location(), initial.Location())
	}
}

func TestTurnLeft(t *testing.T) {
	p := NewPlanet(5, 5)

	assertRobot(t, Robot{p, 1, 1, N}, Robot{p, 1, 1, W}, func(r *Robot) { r.turnLeft() })
	assertRobot(t, Robot{p, 1, 1, E}, Robot{p, 1, 1, N}, func(r *Robot) { r.turnLeft() })
	assertRobot(t, Robot{p, 1, 1, S}, Robot{p, 1, 1, E}, func(r *Robot) { r.turnLeft() })
	assertRobot(t, Robot{p, 1, 1, W}, Robot{p, 1, 1, S}, func(r *Robot) { r.turnLeft() })
}

func TestTurnRight(t *testing.T) {
	p := NewPlanet(5, 5)

	assertRobot(t, Robot{p, 1, 1, N}, Robot{p, 1, 1, E}, func(r *Robot) { r.turnRight() })
	assertRobot(t, Robot{p, 1, 1, E}, Robot{p, 1, 1, S}, func(r *Robot) { r.turnRight() })
	assertRobot(t, Robot{p, 1, 1, S}, Robot{p, 1, 1, W}, func(r *Robot) { r.turnRight() })
	assertRobot(t, Robot{p, 1, 1, W}, Robot{p, 1, 1, N}, func(r *Robot) { r.turnRight() })
}

func TestMove(t *testing.T) {
	p := NewPlanet(5, 5)

	assertRobot(t, Robot{p, 1, 1, N}, Robot{p, 1, 2, N}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 1, 1, E}, Robot{p, 2, 1, E}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 1, 1, S}, Robot{p, 1, 0, S}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 1, 1, W}, Robot{p, 0, 1, W}, func(r *Robot) { r.move() })
}

func TestMoveOutOfBounds(t *testing.T) {
	p := NewPlanet(5, 5)

	assertRobot(t, Robot{p, 0, 0, S}, Robot{p, 0, 0, S}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 0, 0, W}, Robot{p, 0, 0, W}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 5, 5, N}, Robot{p, 5, 5, N}, func(r *Robot) { r.move() })
	assertRobot(t, Robot{p, 5, 5, E}, Robot{p, 5, 5, E}, func(r *Robot) { r.move() })
}
