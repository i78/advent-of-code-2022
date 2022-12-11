package treeGrid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "30373\n25512\n65332\n33549\n35390"

func TestParseGrid(t *testing.T) {
	t.Run("should return grid as expected", func(t *testing.T) {
		result := NewGrid(testdata)

		expectedResult := Grid{
			{Height(3), Height(0), Height(3), Height(7), Height(3)},
			{Height(2), Height(5), Height(5), Height(1), Height(2)},
			{Height(6), Height(5), Height(3), Height(3), Height(2)},
			{Height(3), Height(3), Height(5), Height(4), Height(9)},
			{Height(3), Height(5), Height(3), Height(9), Height(0)},
		}

		assert.Equal(t, expectedResult, result)
	})
}
