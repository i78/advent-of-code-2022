package main

import (
	"dreese.de/aoc22/day12/elevations"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		fmt.Printf("Part 1: %d\nPart 2: %d", SolveA(content), SolveB(content))
	} else {
		log.Fatal(err)
	}
}

func SolveA(content []byte) int {
	elevationMap, start, summit := elevations.NewElevationsMap(string(content))
	dist, _ := elevationMap.DistanceBfsWithCriteria(
		start,
		elevationMap.NodeCoordinateEqualsCriteria(summit),
		elevationMap.UpwardsElevationConstraint())
	return dist
}

func SolveB(content []byte) int {
	elevationMap, _, summit := elevations.NewElevationsMap(string(content))
	dist, _ := elevationMap.DistanceBfsWithCriteria(
		summit,
		elevationMap.AtEqualsCriteria('a'),
		elevationMap.DownwardsElevationConstraint())
	return dist
}
