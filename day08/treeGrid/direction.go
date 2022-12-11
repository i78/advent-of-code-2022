package treeGrid

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (g Grid) Walk(x int, y int, direction Direction) (xNext int, yNext int, exists bool) {
	switch direction {
	case North:
		xNext, yNext = x, y-1
		exists = yNext >= 0
	case South:
		xNext, yNext = x, y+1
		exists = yNext < len(g)
	case East:
		xNext, yNext = x+1, y
		exists = xNext < len(g[y])
	case West:
		xNext, yNext = x-1, y
		exists = xNext >= 0
	}
	return
}
