package main

import (
	"dreese.de/aoc22/day9/rope"
	"os"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	moves := rope.NewMovesList(string(content))

	for i := 0; i < b.N; i++ {
		SolveA(moves)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	moves := rope.NewMovesList(string(content))

	for i := 0; i < b.N; i++ {
		SolveB(moves)
	}
}
