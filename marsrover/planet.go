package marsrover

type Planet struct {
	x int
	y int
}

func NewPlanet(x int, y int) (p *Planet) {
	return &Planet{x, y}
}

func (p *Planet) IsInBounds(x int, y int) bool {
	return x <= p.x && y <= p.y && x >= 0 && y >= 0
}
