package rope

type Move struct {
	Direction
	Steps int
}

type Direction byte

const (
	Left  Direction = 'L'
	Right           = 'R'
	Up              = 'U'
	Down            = 'D'
)

type Moves []Move

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Step(direction Direction) (result Coordinate) {
	switch direction {
	case Up:
		result.X, result.Y = c.X, c.Y+1
	case Down:
		result.X, result.Y = c.X, c.Y-1
	case Left:
		result.X, result.Y = c.X-1, c.Y
	case Right:
		result.X, result.Y = c.X+1, c.Y
	}
	return
}
