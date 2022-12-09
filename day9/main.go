package main

import (
	"dreese.de/aoc22/day9/rope"
	"fmt"
	"github.com/samber/lo"
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
	var visits []rope.Coordinate
	visitInterceptor := func(_ int, coordinate rope.Coordinate) {
		if lo.IndexOf(visits, coordinate) == -1 {
			visits = append(visits, coordinate)
		}
	}

	shortRobe := rope.Rope{
		Head:  rope.Coordinate{X: 0, Y: 0},
		Tails: make([]rope.Coordinate, 1, 1),
	}

	for _, move := range moves {
		shortRobe.MoveHead(move.Direction, move.Steps, visitInterceptor)
	}

	return len(visits)
}

func SolveB(moves rope.Moves) int {
	var visits []rope.Coordinate
	visitorInterceptor := func(tailId int, coordinate rope.Coordinate) {
		if tailId == 8 && lo.IndexOf(visits, coordinate) == -1 {
			visits = append(visits, coordinate)
		}
	}

	longRobe := rope.Rope{
		Head:  rope.Coordinate{X: 0, Y: 0},
		Tails: make([]rope.Coordinate, 9, 9),
	}

	for _, move := range moves {
		longRobe.MoveHead(move.Direction, move.Steps, visitorInterceptor)
	}
	return len(visits)
}
