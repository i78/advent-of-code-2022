package main

import (
	"dreese.de/aoc22/day11/monkey"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {

		fmt.Printf("Part 1: %d\nPart 2: %d", SolveA(content), SolveB(content))
	} else {
		log.Fatal(err)
	}
}

func SolveA(content []byte) int {
	monkeys := monkey.NewMonkeyList(string(content))
	strategy := monkey.DivideStrategy(3)

	for i := 0; i < 20; i++ {
		monkeys = monkeys.NextRound(strategy)
	}

	sort.Slice(monkeys[:], func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	return monkeys[0].InspectionCount * monkeys[1].InspectionCount
}

func SolveB(content []byte) int  {
	monkeys := monkey.NewMonkeyList(string(content))
	strategy := monkey.ModulusStrategy(monkeys)

	for i := 0; i < 10000; i++ {
		monkeys = monkeys.NextRound(strategy)
	}

	sort.Slice(monkeys[:], func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	return monkeys[0].InspectionCount * monkeys[1].InspectionCount
}