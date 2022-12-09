package rope

import (
	"math"
)

type Rope struct {
	Knots []Coordinate
}

type VisitDecorator struct {
	ActOnKnot int
	Fn        func(coordinate Coordinate)
}

func (r *Rope) MoveHead(direction Direction, steps int, decorators ...VisitDecorator) *Rope {
	for step := 0; step < steps; step++ {
		r.Knots[0].Step(direction)

		for tailId := 1; tailId < len(r.Knots); tailId++ {
			current, previous := &r.Knots[tailId], &r.Knots[tailId-1]

			if !current.IsAdjectant(previous) {
				current.EnsureAdjectant(previous)
			}

			for _, i := range decorators {
				if i.ActOnKnot == tailId {
					i.Fn(r.Knots[tailId])
				}
			}
		}
	}
	return r
}

func (c *Coordinate) EnsureAdjectant(other *Coordinate) *Coordinate {
	dx := other.X - c.X
	dy := other.Y - c.Y

	if dx == 0 && math.Abs(float64(dy)) == 2 {
		c.Y += dy / 2
	}

	if math.Abs(float64(dx)) == 2 && dy == 0 {
		c.X += dx / 2
	}

	if math.Abs(float64(dx)) == 2 && math.Abs(float64(dy)) == 1 {
		c.X += dx / 2
		c.Y += dy
	}

	if math.Abs(float64(dx)) == 1 && math.Abs(float64(dy)) == 2 {
		c.X += dx
		c.Y += dy / 2
	}

	if math.Abs(float64(dx)) == 2 && math.Abs(float64(dy)) == 2 {
		c.X += dx / 2
		c.Y += dy / 2
	}

	if dx > 2 || dy > 2 {
		panic("rope was torn!")
	}

	return c
}

func (c *Coordinate) IsAdjectant(other *Coordinate) bool {
	return math.Abs(float64(c.X-other.X)) <= 1 && math.Abs(float64(c.Y-other.Y)) <= 1
}
