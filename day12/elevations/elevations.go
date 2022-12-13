package elevations

import "C"

type ElevationMap [][]byte

type ComparisonCriteria = func(coordinate Coordinate) bool
type NextNodeConstraint func(current Coordinate, next Coordinate) bool

func (m *ElevationMap) NodeCoordinateEqualsCriteria(destination Coordinate) ComparisonCriteria {
	return func(coordinate Coordinate) bool {
		return coordinate == destination
	}
}

func (m *ElevationMap) AtEqualsCriteria(value byte) ComparisonCriteria {
	return func(coordinate Coordinate) bool {
		return m.At(coordinate) == value
	}
}

func (m *ElevationMap) UpwardsElevationConstraint() NextNodeConstraint {
	return func(current Coordinate, next Coordinate) bool {
		return m.gradient(current, next) <= 1
	}
}

func (m *ElevationMap) DownwardsElevationConstraint() NextNodeConstraint {
	return func(current Coordinate, next Coordinate) bool {
		return m.gradient(current, next) >= -1
	}
}

func (m *ElevationMap) DistanceBfsWithCriteria(start Coordinate, foundCriteria ComparisonCriteria, elevationConstraint NextNodeConstraint) (int, bool) {
	type Explored struct {
		parent *Coordinate
	}

	explored := make(map[Coordinate]Explored)
	queue := []Coordinate{start}
	explored[start] = Explored{parent: nil}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if foundCriteria(vertex) {
			steps := -1
			for vert := &vertex; vert != nil; steps++ {
				exploredRecord, _ := explored[*vert]
				vert = exploredRecord.parent
			}
			return steps, true
		}

		for _, direction := range []Direction{East, South, West, North} {
			next := vertex.Step(direction)

			if m.Exists(next) && elevationConstraint(vertex, next) {
				if _, alreadyExplored := explored[next]; !alreadyExplored {
					explored[next] = Explored{parent: &vertex}
					queue = append(queue, next)
				}
			}
		}
	}
	return 0, false
}

func (m *ElevationMap) gradient(v Coordinate, next Coordinate) int {
	return int(m.At(next)) - int(m.At(v))
}

func (m *ElevationMap) Exists(coordinate Coordinate) bool {
	return coordinate.X >= 0 &&
		coordinate.Y >= 0 &&
		coordinate.Y < len(*m) &&
		coordinate.X < len((*m)[coordinate.Y])
}

func (m *ElevationMap) At(coordinate Coordinate) byte {
	return (*m)[coordinate.Y][coordinate.X]
}
