package cave

import (
	"fmt"
	"github.com/samber/lo"
	"math"
)

const (
	Rock = '#'
	Sand = 'o'
)

type Cell byte
type Cave map[Coordinate]Cell

func (c *Cave) Print() {
	for y := 0; y <= 150; y++ {
		for x := 460; x <= 550; x++ {
			if at, exists := (*c)[Coordinate{x, y}]; exists {
				fmt.Print(lo.Ternary(at == Sand,
					"\033[33mo",
					"\033[37m▆"))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (c *Cave) IsOccupied(coordinate Coordinate, limits Limits) bool {
	_, occupied := (*c)[coordinate]
	return coordinate.Y >= limits.FloorHeight || occupied
}

func (c *Cave) Drop(limits Limits) (landedAt Coordinate, complete bool) {
	sand := Coordinate{500, 0}
	for current, next := sand, sand; ; current, next = next, next.Step(Down) {

		if c.IsOccupied(next, limits) {
			if leftDown := current.Step(LeftDown); !c.IsOccupied(leftDown, limits) {
				next = leftDown
			} else if rightDown := current.Step(RightDown); !c.IsOccupied(rightDown, limits) {
				next = rightDown
			} else if current != sand {
				(*c)[current] = Sand
				return current, false
			} else {
				return next, true
			}
		}

		if next.Y > limits.StartOfVoid {
			return next, true
		}
	}
}

type Limits struct {
	StartOfVoid int
	FloorHeight int
}

func (l *Limits) ToFloor() Limits {
	return Limits{
		StartOfVoid: math.MaxInt,
		FloorHeight: l.StartOfVoid + 2,
	}
}
