package cave

type Direction string

const (
	Left      Direction = "W"
	Right               = "E"
	Up                  = "N"
	Down                = "S"
	LeftDown            = "SW"
	RightDown           = "SE"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Step(direction Direction) (coordinate Coordinate) {
	switch direction {
	case Up:
		coordinate = Coordinate{c.X, c.Y - 1}
	case Down:
		coordinate = Coordinate{c.X, c.Y + 1}
	case Left:
		coordinate = Coordinate{c.X - 1, c.Y}
	case Right:
		coordinate = Coordinate{c.X + 1, c.Y}
	case LeftDown:
		coordinate = c.Step(Down).Step(Left)
	case RightDown:
		coordinate = c.Step(Down).Step(Right)
	}

	return
}
