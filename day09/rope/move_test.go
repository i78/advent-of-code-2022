package rope

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveHead(t *testing.T) {

	t.Run("Should move to expected coordinates", func(t *testing.T) {
		testcases := []struct {
			start Coordinate
			Direction
			steps    int
			expected Coordinate
		}{
			{
				start:     Coordinate{X: 0, Y: 0},
				Direction: Left,
				expected:  Coordinate{X: -1, Y: 0},
			}, {
				start:     Coordinate{X: 10, Y: 10},
				Direction: Right,
				expected:  Coordinate{X: 11, Y: 10},
			},
		}

		for _, testcase := range testcases {
			t.Run(fmt.Sprintf("Should return %v for %v,%c", testcase.expected, testcase.start, testcase.Direction), func(t *testing.T) {
				testcase.start.Step(testcase.Direction)
				assert.Equal(t, testcase.expected, testcase.start)
			})
		}

	})
}
