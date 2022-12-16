package cave

import (
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
)

const (
	LineSeparator = "\n"
	Arrow         = " -> "
)

func parseCoordiante(s string) Coordinate {
	tokens := strings.Split(s, ",") //cons
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return Coordinate{x, y}
}

func NewCave(raw string) (Cave, Limits) {
	result := make(Cave)
	limits := Limits{
		StartOfVoid: -1,
		FloorHeight: math.MaxInt,
	}

	for _, pair := range strings.Split(raw, LineSeparator) {
		snaps := strings.Split(pair, Arrow)

		for idx := range snaps[:len(snaps)-1] {
			start := parseCoordiante(snaps[idx])
			end := parseCoordiante(snaps[idx+1])

			if start.X != end.X && start.Y != end.Y {
				panic("Corrupted input data")
			}

			for x, y := start.X, start.Y; x != end.X || y != end.Y; {
				result[Coordinate{x, y}] = Rock

				switch {
				case x < end.X:
					x++
				case x > end.X:
					x--
				case y < end.Y:
					y++
				case y > end.Y:
					y--
				default:
				}

				limits.StartOfVoid = lo.Max([]int{y, limits.StartOfVoid})

				result[Coordinate{x, y}] = Rock
			}
		}
	}

	return result, limits
}
