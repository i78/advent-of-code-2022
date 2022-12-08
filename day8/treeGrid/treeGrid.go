package treeGrid

type Height int8
type Row []Height
type Grid [][]Height

func (g Grid) IsOutsideEdge(x int, y int) bool {
	return x == 0 || y == 0 || y == len(g)-1 || x == len(g[x])-1
}

func (g Grid) IsVisibleFromOutside(x int, y int) bool {
	if g.IsOutsideEdge(x, y) {
		return true
	}

	for _, direction := range []Direction{North, East, South, West} {
		max := Height(0)
		for xNext, yNext, exists := g.Walk(x, y, direction); exists; xNext, yNext, exists = g.Walk(xNext, yNext, direction) {
			if g[yNext][xNext] > max {
				max = g[yNext][xNext]
			}
			if max > g[y][x] {
				break
			}
		}
		if max < g[y][x] {
			return true
		}
	}

	return false
}

func (g Grid) ScenicScore(x int, y int) (score int) {
	if g.IsOutsideEdge(x, y) {
		return 0
	}
	score = 1

	for _, direction := range []Direction{North, East, South, West} {
		viewWidth := 0
		for xNext, yNext, exists := g.Walk(x, y, direction); exists; xNext, yNext, exists = g.Walk(xNext, yNext, direction) {
			viewWidth++
			if g[yNext][xNext] >= g[y][x] {
				break
			}
		}
		score *= viewWidth
	}
	return
}
