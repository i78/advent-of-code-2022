package rope

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAdjectant(t *testing.T) {
	testcases := []struct {
		a        Coordinate
		b        Coordinate
		expected bool
	}{
		{
			a:        Coordinate{X: 0, Y: 0},
			b:        Coordinate{X: -1, Y: 0},
			expected: true,
		}, {
			a:        Coordinate{X: 10, Y: 10},
			b:        Coordinate{X: 11, Y: 0},
			expected: false,
		}, {
			a:        Coordinate{X: 10, Y: 10},
			b:        Coordinate{X: 10, Y: 8},
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %t for %v,%v", testcase.expected, testcase.a, testcase.b), func(t *testing.T) {
			assert.Equal(t, testcase.expected, testcase.a.IsAdjectant(&testcase.b))
		})
	}
}

func TestEnsureAdjectant(t *testing.T) {
	testcases := []struct {
		a        Coordinate
		b        Coordinate
		expected Coordinate
	}{
		{
			a:        Coordinate{X: 0, Y: 0},
			b:        Coordinate{X: -1, Y: 0},
			expected: Coordinate{X: 0, Y: 0},
		}, {
			a:        Coordinate{X: 10, Y: 0},
			b:        Coordinate{X: 12, Y: 0},
			expected: Coordinate{X: 11, Y: 0},
		}, {
			a:        Coordinate{X: 0, Y: 0},
			b:        Coordinate{X: 2, Y: 2},
			expected: Coordinate{X: 1, Y: 1},
		}, {
			a:        Coordinate{X: 2, Y: 0},
			b:        Coordinate{X: 4, Y: 1},
			expected: Coordinate{X: 3, Y: 1},
		}, {
			a:        Coordinate{X: 5, Y: 4},
			b:        Coordinate{X: 3, Y: 5},
			expected: Coordinate{X: 4, Y: 5},
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %v for %v,%v", testcase.expected, testcase.a, testcase.b), func(t *testing.T) {
			testcase.a.EnsureAdjectant(&testcase.b)
			assert.Equal(t, testcase.expected, testcase.a)
		})
	}
}

func TestShouldRemainAdjectant(t *testing.T) {
	testcases := []struct {
		before Rope
		Direction
		steps    int
		expected Rope
	}{
		{
			before: Rope{
				Head:  Coordinate{X: 0, Y: 0},
				Tails: []Coordinate{{X: 0, Y: 0}},
			},
			Direction: Right,
			steps:     2,
			expected: Rope{
				Head:  Coordinate{X: 2, Y: 0},
				Tails: []Coordinate{{X: 1, Y: 0}},
			},
		}, {
			before: Rope{
				Head:  Coordinate{X: 0, Y: 0},
				Tails: []Coordinate{{X: 0, Y: 0}},
			},
			Direction: Down,
			steps:     2,
			expected: Rope{
				Head:  Coordinate{X: 0, Y: -2},
				Tails: []Coordinate{{X: 0, Y: -1}},
			},
		}, {
			before: Rope{
				Head:  Coordinate{X: 4, Y: 1},
				Tails: []Coordinate{{X: 3, Y: 0}},
			},
			Direction: Up,
			steps:     1,
			expected: Rope{
				Head:  Coordinate{X: 4, Y: 2},
				Tails: []Coordinate{{X: 4, Y: 1}},
			},
		}, {
			before: Rope{
				Head:  Coordinate{X: 4, Y: 1},
				Tails: []Coordinate{{X: 3, Y: 0}, {X: 2, Y: 0}},
			},
			Direction: Up,
			steps:     2,
			expected: Rope{
				Head:  Coordinate{X: 4, Y: 3},
				Tails: []Coordinate{{X: 4, Y: 2}, {X: 3, Y: 1}},
			},
		},
	}

	for idx, testcase := range testcases {
		t.Run(fmt.Sprintf("%d. Should return %v for %v", idx, testcase.before, testcase.expected), func(t *testing.T) {
			subject := testcase.before.MoveHead(testcase.Direction, testcase.steps)
			assert.Equal(t, &testcase.expected, subject)
		})
	}
}

func TestRecordVisits(t *testing.T) {
	testcases := []struct {
		Moves
		Rope
		observedTail        int
		expectedNumOfVisits int
	}{
		{
			Moves: Moves{
				{Right, 4},
				{Up, 4},
				{Left, 3},
				{Down, 1},
				{Right, 4},
				{Down, 1},
				{Left, 5},
				{Right, 2},
			},
			Rope: Rope{
				Head:  Coordinate{X: 0, Y: 0},
				Tails: []Coordinate{{X: 0, Y: 0}},
			},
			observedTail:        0,
			expectedNumOfVisits: 13,
		}, {
			Moves: Moves{
				{Right, 5},
				{Up, 8},
				{Left, 8},
				{Down, 3},
				{Right, 17},
				{Down, 10},
				{Left, 25},
				{Up, 20},
			},
			Rope: Rope{
				Head:  Coordinate{X: 0, Y: 0},
				Tails: make([]Coordinate, 9, 9),
			},
			observedTail:        8,
			expectedNumOfVisits: 36,
		},
	}

	for caseId, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %d for %d", testcase.expectedNumOfVisits, caseId), func(t *testing.T) {
			var visits []Coordinate
			visitInterceptor := func(tailId int, coordinate Coordinate) {
				if tailId == testcase.observedTail && lo.IndexOf(visits, coordinate) == -1 {
					visits = append(visits, coordinate)
				}
			}

			for _, move := range testcase.Moves {
				testcase.Rope.MoveHead(move.Direction, move.Steps, visitInterceptor)
			}

			assert.Len(t, visits, testcase.expectedNumOfVisits)

		})
	}
}

func TestRobeWithMultipleKnots(t *testing.T) {
	fakeRobe := func() Rope {
		return Rope{
			Head:  Coordinate{X: 0, Y: 0},
			Tails: []Coordinate{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}},
		}
	}

	t.Run("Should move all tails out of start when moved away 4 steps", func(t *testing.T) {
		subject := fakeRobe()
		subject.MoveHead(Right, 4)
		assert.Equal(t, []Coordinate{{X: 3, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 0}}, subject.Tails)
	})

	t.Run("Should move all tails out of start and 90deg shift when moved away 4,4 steps", func(t *testing.T) {
		subject := fakeRobe()
		subject.MoveHead(Right, 4).
			MoveHead(Up, 6)

		assert.Equal(t, []Coordinate{{X: 4, Y: 5}, {X: 4, Y: 4}, {X: 4, Y: 3}}, subject.Tails)
	})
}
