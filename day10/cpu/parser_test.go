package cpu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "noop\naddx 3\naddx -5"

func TestParseGrid(t *testing.T) {
	t.Run("should return grid as expected", func(t *testing.T) {
		result := NewProgram(testdata)

		expectedResult := Program{
			Noop{},
			Addx{3},
			Addx{-5},
		}

		assert.Equal(t, expectedResult, result)
	})
}
