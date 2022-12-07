package main

import (
	"dreese.de/aoc22/day7/fstree"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		tree := fstree.NewFsTree(string(content))
		allDirs := tree.FindAllDirectories()

		// Part A
		const maxDirSizeToBeConsidered = 100000
		totalSmallDirectorySizes := totalSizeOfSmallDirectories(allDirs, maxDirSizeToBeConsidered)
		fmt.Printf("Part 1: %d\n", totalSmallDirectorySizes)

		// Part B
		const size, required = 70000000, 30000000
		requiredToFree := spaceToBeReclaimed(size, required, tree.DirectorySize())
		smallestViableDirectorySize := smallestDirectorySuitableForDeletion(allDirs, requiredToFree)

		fmt.Printf("Part 2: %d\n", smallestViableDirectorySize)
	} else {
		log.Fatal(err)
	}
}

func totalSizeOfSmallDirectories(directories []*fstree.Node, maxSize int) int {
	return lo.Sum(
		lo.Filter(
			lo.Map(directories, func(it *fstree.Node, _ int) int { return it.DirectorySize() }),
			func(it int, _ int) bool { return it <= maxSize }))
}

func spaceToBeReclaimed(size int, required int, used int) int {
	return required - (size - used)
}

func smallestDirectorySuitableForDeletion(directories []*fstree.Node, spaceToReclaim int) int {
	return lo.Min(
		lo.Filter(
			lo.Map(directories, func(it *fstree.Node, _ int) int { return it.DirectorySize() }),
			func(it int, _ int) bool { return it >= spaceToReclaim }))
}
