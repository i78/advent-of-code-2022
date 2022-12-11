package main

import (
	"dreese.de/aoc22/day3/rucksack"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		r := rucksack.NewRucksackList(string(content))

		commonItems := r.CommonItems()
		totalScore := rucksack.TotalScore(commonItems)

		fmt.Printf("Part 1 result = %d \n", totalScore)

		// Part B
		groups := r.ToGroups(3)
		totalItems := rucksack.AllBadgeItemsFrom(groups)
		score := rucksack.TotalScore(totalItems)
		fmt.Printf("Part 2 result = %d \n", score)

	} else {
		log.Fatal(err)
	}
}
