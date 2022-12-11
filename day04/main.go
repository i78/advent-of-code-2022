package main

import (
	"dreese.de/aoc22/day3/sections"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		// Part A
		data := sections.NewSectionList(string(content))

		fullOverlapsCount := lo.CountBy[sections.SectionRangePair](data, func(it sections.SectionRangePair) bool {
			return it.Overlaps(sections.FullyIncludes)
		})

		fmt.Printf("Part 1 result = %d \n", fullOverlapsCount)

		// Part B
		numOfOverlaps := lo.CountBy[sections.SectionRangePair](data, func(it sections.SectionRangePair) bool {
			return it.Overlaps(sections.Overlaps)
		})

		fmt.Printf("Part 2 result = %d \n", numOfOverlaps)
	} else {
		log.Fatal(err)
	}
}
