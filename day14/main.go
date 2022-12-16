package main

import (
	"dreese.de/aoc22/day14/cave"
	"fmt"
	escapes "github.com/snugfox/ansi-escapes"
	"log"
	"os"
	"time"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		_, fancy := os.LookupEnv("FANCY")
		resultA, resultB := SolveA(content, fancy), SolveB(content, fancy)
		fmt.Printf("Part 1: %d\nPart 2: %d", resultA, resultB)
	} else {
		log.Fatal(err)
	}
}

func SolveA(content []byte, fancy bool) int {
	caveWithAbyss, limits := cave.NewCave(string(content))
	for i := 0; i < 6553; i++ {
		_, droppedIntoAbyss := caveWithAbyss.Drop(limits)

		if droppedIntoAbyss == true {
			return i
		}

		if fancy {
			fmt.Print(escapes.CursorPos(1, 1))
			caveWithAbyss.Print()
			time.Sleep(1 * time.Millisecond)
		}

	}
	return -1
}

func SolveB(content []byte, fancy bool) int {
	caveWithFloor, limits := cave.NewCave(string(content))
	limits = limits.ToFloor()
	for i := 0; i < 65535; i++ {
		_, caveIsFull := caveWithFloor.Drop(limits)

		if caveIsFull == true {
			return i + 1
		}

		if fancy {
			fmt.Print(escapes.CursorPos(1, 1))
			caveWithFloor.Print()
			time.Sleep(1 * time.Millisecond)
		}

	}
	return -1
}
