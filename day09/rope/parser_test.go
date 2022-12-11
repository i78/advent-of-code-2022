package rope

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2"

func TestParseGrid(t *testing.T) {
	t.Run("should return grid as expected", func(t *testing.T) {
		result := NewMovesList(testdata)

		expectedResult := Moves{
			{Right, 4},
			{Up, 4},
			{Left, 3},
			{Down, 1},
			{Right, 4},
			{Down, 1},
			{Left, 5},
			{Right, 2},
		}

		assert.Equal(t, expectedResult, result)
	})
}
