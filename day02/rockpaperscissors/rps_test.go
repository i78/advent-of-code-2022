package rockpaperscissors

import (
	"fmt"
	"testing"
)
import assert "github.com/stretchr/testify/assert"

const testdata = "A Y\nB X\nC Z"

func TestWhenTestDataParsed(t *testing.T) {

	t.Run("Should return result with expected count of 3 (rounds)", func(t *testing.T) {
		subject := NewRoundsList(testdata)

		const expectedNumberOfRounds = 3
		assert.Len(t, subject, expectedNumberOfRounds)
	})

	t.Run("Should parse to expected rounds data", func(t *testing.T) {
		expectedRounds := []Round{
			{Opponent: Rock, Player: Paper},
			{Opponent: Paper, Player: Rock},
			{Opponent: Scissors, Player: Scissors},
		}

		subject := NewRoundsList(testdata)

		assert.Equal(t, expectedRounds, subject)
	})

	t.Run("should parse to expected rounds data in RequiredResultsJob format", func(t *testing.T) {
		expected := RequiredResultsJobs{
			{Round: Round{Opponent: Rock}, RequiredResult: Draw},
			{Round: Round{Opponent: Paper}, RequiredResult: Lose},
			{Round: Round{Opponent: Scissors}, RequiredResult: Win},
		}

		subject := NewRequiredScoreJobList(testdata)

		assert.Equal(t, expected, subject)
	})
}

func TestMatchResults(t *testing.T) {
	testcases := []struct {
		input         Round
		expectedScore int
	}{
		{input: Round{Opponent: Rock, Player: Paper}, expectedScore: 8},
		{input: Round{Opponent: Paper, Player: Rock}, expectedScore: 1},
		{input: Round{Opponent: Scissors, Player: Scissors}, expectedScore: 6},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return expected result %d", testcase.expectedScore), func(t *testing.T) {
			result := testcase.input.Score()
			assert.Equal(t, testcase.expectedScore, result)
		})
	}
}

func TestTotalScore(t *testing.T) {
	rounds := []Round{
		{Opponent: Rock, Player: Paper},
		{Opponent: Paper, Player: Rock},
		{Opponent: Scissors, Player: Scissors},
	}

	const expectedScore = 15

	t.Run("should return expected score", func(t *testing.T) {
		score := TotalScore(rounds)

		assert.Equal(t, expectedScore, score)
	})
}

func TestExpectedScore(t *testing.T) {
	input := RequiredResultsJobs{
		{Round: Round{Opponent: Rock}, RequiredResult: Draw},
		{Round: Round{Opponent: Paper}, RequiredResult: Lose},
		{Round: Round{Opponent: Scissors}, RequiredResult: Win},
	}

	t.Run("should map input to expected rounds", func(t *testing.T) {
		expectedRounds := []Round{
			{Opponent: Rock, Player: Rock},
			{Opponent: Paper, Player: Rock},
			{Opponent: Scissors, Player: Rock},
		}

		var subject = input.MapToRounds()
		assert.Equal(t, expectedRounds, subject)
	})

	t.Run("should have expected score", func(t *testing.T) {
		var subject = input.MapToRounds()

		const expectedScore = 12
		assert.Equal(t, expectedScore, TotalScore(subject))
	})

}
