package rockpaperscissors

import (
	"strings"
)

type Rps int

const (
	Rock     Rps = 1
	Paper        = 2
	Scissors     = 3
)

type MatchResult int

const (
	Draw MatchResult = iota
	Lose
	Win
)

type Round struct {
	Opponent Rps
	Player   Rps
}

type RequiredResultsJob struct {
	Round
	RequiredResult MatchResult
}

type RequiredResultsJobs []RequiredResultsJob

const TelegramSeparator = "\n"
const ItemSeparator = " "

func MapRecord(record string) (result Rps) {
	switch strings.TrimSpace(record) {
	case "A", "X":
		result = Rock
	case "B", "Y":
		result = Paper
	case "C", "Z":
		result = Scissors
	default:
		panic("BOOM")
	}
	return
}

func MapRequiredResult(s string) (result MatchResult) {
	switch strings.TrimSpace(s) {
	case "X":
		result = Lose
	case "Y":
		result = Draw
	case "Z":
		result = Win
	default:
		panic("BOOM")
	}
	return
}

func NewRoundsList(raw string) (result []Round) {
	for _, telegram := range strings.Split(raw, TelegramSeparator) {
		records := strings.Split(telegram, ItemSeparator)
		opponent := MapRecord(records[0])
		player := MapRecord(records[1])

		result = append(result, Round{
			Opponent: opponent,
			Player:   player,
		})
	}
	return
}

func NewRequiredScoreJobList(raw string) (result RequiredResultsJobs) {
	for _, telegram := range strings.Split(raw, TelegramSeparator) {
		records := strings.Split(telegram, ItemSeparator)
		opponent := MapRecord(records[0])
		requiredResult := MapRequiredResult(records[1])

		result = append(result, RequiredResultsJob{
			Round{
				Opponent: opponent,
			},
			requiredResult,
		})
	}
	return
}

func (jobs *RequiredResultsJobs) MapToRounds() (result []Round) {
	for _, job := range *jobs {
		choices := []Rps{Rock, Paper, Scissors}

		with := func(playerChoice Rps) Round {
			return Round{
				Opponent: job.Opponent,
				Player:   playerChoice,
			}
		}

		for _, choice := range choices {
			if attempt := with(choice); attempt.MatchResult() == job.RequiredResult {
				result = append(result, attempt)
			}
		}
	}
	return
}

func (r *Round) MatchResult() MatchResult {
	if r.Player == r.Opponent {
		return Draw
	}

	beats := map[Rps]Rps{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}

	if beats[r.Player] == r.Opponent {
		return Win
	}

	return Lose
}

func (r *Round) Score() (score int) {
	switch r.MatchResult() {
	case Win:
		score += 6
	case Draw:
		score += 3
	}

	score += int(r.Player)

	return score
}

func TotalScore(rounds []Round) (score int) {
	for _, round := range rounds {
		score += round.Score()
	}
	return
}
