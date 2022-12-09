package rope

import (
	"math"
)

type Rope struct {
	Head  Coordinate
	Tails []Coordinate
}

type VisitInterceptor func(tailId int, coordinate Coordinate)

func (r *Rope) MoveHead(direction Direction, steps int, interceptors ...VisitInterceptor) *Rope {
	for step := 0; step < steps; step++ {
		r.Head = r.Head.Step(direction)

		for tailId := range r.Tails {
			current := &r.Tails[tailId]
			var previous *Coordinate

			if tailId == 0 {
				previous = &r.Head
			} else {
				previous = &r.Tails[tailId-1]
			}

			if !current.IsAdjectant(previous) {
				current.EnsureAdjectant(previous)
			}

			for _, i := range interceptors {
				i(tailId, r.Tails[tailId])
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
