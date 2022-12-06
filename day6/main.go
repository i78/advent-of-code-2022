package main

import (
	"fmt"
	"log"
	"os"
)

const StxChunkWidth = 4
const MsgChunkWidth = 14

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		// Part A
		fmt.Printf("Part 1: %d\n", FirstDistinct(StxChunkWidth, content))

		// Part B
		fmt.Printf("Part 2: %d\n", FirstDistinct(MsgChunkWidth, content))

	} else {
		log.Fatal(err)
	}
}
