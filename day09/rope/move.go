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
	X int32
	Y int32
}

func (c *Coordinate) Step(direction Direction) {
	switch direction {
	case Up:
		c.Y = c.Y + 1
	case Down:
		c.Y = c.Y - 1
	case Left:
		c.X = c.X - 1
	case Right:
		c.X = c.X + 1
	}
}

func (m Moves) Apply(rope Rope, interceptors ...VisitDecorator) {
	for _, move := range m {
		rope.MoveHead(move.Direction, move.Steps, interceptors...)
	}
}
