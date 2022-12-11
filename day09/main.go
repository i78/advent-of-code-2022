package main

import (
	"dreese.de/aoc22/day9/rope"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		moves := rope.NewMovesList(string(content))

		fmt.Printf("Part 1: %d\nPart 2: %d\n", SolveA(moves), SolveB(moves))
	} else {
		log.Fatal(err)
	}
}

func SolveA(moves rope.Moves) int {
	visits, visitCountingDecorator := buildVisitCountingDecorator(1)

	shortRope := rope.Rope{
		Knots: make([]rope.Coordinate, 2),
	}

	moves.Apply(shortRope, visitCountingDecorator)

	return len(*visits)
}

func SolveB(moves rope.Moves) int {
	visits, visitCountingDecorator := buildVisitCountingDecorator(9)

	longRobe := rope.Rope{
		Knots: make([]rope.Coordinate, 10),
	}

	moves.Apply(longRobe, visitCountingDecorator)
	return len(*visits)
}

func buildVisitCountingDecorator(actOnKnot int) (*map[rope.Coordinate]bool, rope.VisitDecorator) {
	visits := make(map[rope.Coordinate]bool, 8192)
	decorator := rope.VisitDecorator{
		ActOnKnot: actOnKnot,
		Fn: func(coordinate rope.Coordinate) {
			visits[coordinate] = true
		},
	}
	return &visits, decorator
}
