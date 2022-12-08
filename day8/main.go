package main

import (
	"dreese.de/aoc22/day8/treeGrid"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		tree := treeGrid.NewGrid(string(content))

		// Part A
		visibleTrees := VisibleTreesParallel(tree)
		fmt.Printf("Part 1: %d\n", visibleTrees)

		// Part B
		maxScenicScore := MaxScenicScoreParallel(tree)
		fmt.Printf("Part 2: %d\n", maxScenicScore)
	} else {
		log.Fatal(err)
	}
}

func MaxScenicScore(tree treeGrid.Grid) int {
	maxScenicScore := 0
	for y, row := range tree {
		for x := range row {
			if score := tree.ScenicScore(x, y); score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	return maxScenicScore
}

func MaxScenicScoreParallel(tree treeGrid.Grid) (maxScore int) {
	resultChan := make(chan int, len(tree))

	for y, row := range tree {
		row := row
		go func(y int, out chan int) {
			rowMaxScore := 0
			for x := range row {
				if score := tree.ScenicScore(x, y); score > rowMaxScore {
					rowMaxScore = score
				}
			}
			resultChan <- rowMaxScore
		}(y, resultChan)
	}

	for i := 0; i < len(tree); i++ {
		if score := <-resultChan; score > maxScore {
			maxScore = score
		}
	}

	return
}

func VisibleTrees(tree treeGrid.Grid) (visible int) {
	for y, row := range tree {
		for x := range row {
			if tree.IsVisibleFromOutside(x, y) {
				visible++
			}
		}
	}
	return visible
}

func VisibleTreesParallel(tree treeGrid.Grid) (visible int) {
	resultChan := make(chan int, len(tree))

	for y, row := range tree {
		row := row
		go func(y int, out chan int) {
			rowVisible := 0
			for x := range row {
				if tree.IsVisibleFromOutside(x, y) {
					rowVisible++
				}
			}
			resultChan <- rowVisible
		}(y, resultChan)
	}

	for i := 0; i < len(tree); i++ {
		visible += <-resultChan
	}

	return visible
}
