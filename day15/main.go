package main

import (
	"dreese.de/aoc22/day15/beacons"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		parser := participle.MustBuild[beacons.ReadingList](
			participle.Lexer(beacons.SensorReadingListLexer),
		)
		readings, _ := parser.ParseString("", string(content))

		resultA, resultB := SolveA(readings), SolveB(readings)

		fmt.Printf("Part 1: %d\nPart 2: %d", resultA, resultB)
	} else {
		log.Fatal(err)
	}
}

func SolveA(content *beacons.ReadingList) int {
	coveredPositions, _ := content.CoveredPositions(2000000)
	return coveredPositions
}

func SolveB(content *beacons.ReadingList) (frequency int) {
	for y := 0; y < 4000000; y++ { // y 5716881
		_, gaps := content.CoveredPositions(y)
		if len(gaps) > 0 {
			frequency = gaps[0]*4000000 + y
			break
		}
	}
	return
}
