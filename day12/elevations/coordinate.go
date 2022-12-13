package elevations

type Direction byte

const (
	West  Direction = 'W'
	East            = 'E'
	North           = 'N'
	South           = 'S'
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Step(direction Direction) (coordinate Coordinate) {
	switch direction {
	case North:
		coordinate = Coordinate{c.X, c.Y - 1}
	case South:
		coordinate = Coordinate{c.X, c.Y + 1}
	case West:
		coordinate = Coordinate{c.X - 1, c.Y}
	case East:
		coordinate = Coordinate{c.X + 1, c.Y}
	}
	return
}
