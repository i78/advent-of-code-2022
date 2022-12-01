package main

import (
	"dreese.de/aoc22/day1/calories"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		caloryRecords := calories.NewCaloryAmounts(string(content))
		maxCalories := calories.FindMaxCaloryValueCarriedByElv(caloryRecords)
		fmt.Printf("Part 1 result = %d \n", maxCalories)

		// Part 2
		top3Calories := calories.CaloriesCarriedByTop3Elves(caloryRecords)
		fmt.Printf("Part 2 result = %d \n", top3Calories)
	} else {
		log.Fatal(err)
	}
}
