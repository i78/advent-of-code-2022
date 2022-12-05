package main

import (
	stacks2 "dreese.de/aoc22/day5/stacks"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func revealResult(stacks stacks2.Stacks) string {
	sb := strings.Builder{}
	for stackId := range stacks {
		top, _ := stacks.PeekTop(stackId)
		sb.WriteByte(top)
	}
	return sb.String()
}

func printStep(stacks *stacks2.Stacks) {
	fmt.Print("\033[H\033[2J")
	stacks.Print()
	time.Sleep(5 * time.Millisecond)
}

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		// Fancy
		stacks, program := stacks2.NewStacks(string(content))
		program.Run(stacks, stacks2.CrateMover9000, printStep)

		// Part A
		stacks, program = stacks2.NewStacks(string(content))
		program.Run(stacks, stacks2.CrateMover9000)
		fmt.Printf("\nPart 1 result=%s", revealResult(stacks))

		// Part B
		stacks, program = stacks2.NewStacks(string(content))
		program.Run(stacks, stacks2.CrateMover9001)
		fmt.Printf("\nPart 2 result=%s", revealResult(stacks))
	} else {
		log.Fatal(err)
	}
}
