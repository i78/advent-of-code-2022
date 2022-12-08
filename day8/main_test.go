package main

import (
	"dreese.de/aoc22/day8/treeGrid"
	"os"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	tree := treeGrid.NewGrid(string(content))

	for i := 0; i < b.N; i++ {
		VisibleTrees(tree)
	}
}

func BenchmarkPartOneParallel(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	tree := treeGrid.NewGrid(string(content))

	for i := 0; i < b.N; i++ {
		VisibleTreesParallel(tree)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	tree := treeGrid.NewGrid(string(content))

	for i := 0; i < b.N; i++ {
		MaxScenicScore(tree)
	}
}

func BenchmarkPartTwoParallel(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	tree := treeGrid.NewGrid(string(content))

	for i := 0; i < b.N; i++ {
		MaxScenicScoreParallel(tree)
	}
}

func BenchmarkPartsCombined(b *testing.B) {
	content, _ := os.ReadFile("input.txt")
	tree := treeGrid.NewGrid(string(content))

	for i := 0; i < b.N; i++ {
		VisibleTreesParallel(tree)
		MaxScenicScoreParallel(tree)
	}
}
