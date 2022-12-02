package main

import (
	"dreese.de/aoc22/day2/rockpaperscissors"
	"fmt"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		rounds := rockpaperscissors.NewRoundsList(string(content))
		totalScore := rockpaperscissors.TotalScore(rounds)

		fmt.Printf("Part 1 result = %d \n", totalScore)

		// Part B
		requiredResults := rockpaperscissors.NewRequiredScoreJobList(string(content))
		roundsMatchingResult := requiredResults.MapToRounds()

		totalScore2 := rockpaperscissors.TotalScore(roundsMatchingResult)

		fmt.Printf("Part 2 result = %d \n", totalScore2)
	} else {
		log.Fatal(err)
	}
}
