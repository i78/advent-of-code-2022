package main

import (
	"dreese.de/aoc22/day10/cpu"
	"dreese.de/aoc22/day10/vic"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		program := cpu.NewProgram(string(content))
		result := cpu.NewCpu().Simulate(program)

		fmt.Printf("Part 1: %d\nPart 2: %s", SolveA(result), SolveB(result))
	} else {
		log.Fatal(err)
	}
}

func SolveA(result cpu.SimulationResult) int {
	cycles := []int{20, 60, 100, 140, 180, 220}

	return lo.SumBy(cycles, func(cycle int) int {
		return result.CpuStateBeforeCycle(cycle).X * cycle
	})
}

func SolveB(result cpu.SimulationResult) string {
	const crtWidth = 40
	return vic.SimulateVic(result, crtWidth)
}
