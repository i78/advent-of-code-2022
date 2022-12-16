package cave

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"

func TestParser(t *testing.T) {
	result, limits := NewCave(testdata)

	t.Run("should parse to expected result", func(t *testing.T) {
		expected := Cave{
			Coordinate{498, 4}: Rock,
			Coordinate{498, 5}: Rock,
			Coordinate{498, 6}: Rock,

			Coordinate{497, 6}: Rock,
			Coordinate{496, 6}: Rock,

			Coordinate{503, 4}: Rock,
			Coordinate{502, 4}: Rock,

			Coordinate{502, 5}: Rock,
			Coordinate{502, 5}: Rock,
			Coordinate{502, 6}: Rock,
			Coordinate{502, 7}: Rock,
			Coordinate{502, 8}: Rock,
			Coordinate{502, 9}: Rock,

			Coordinate{501, 9}: Rock,
			Coordinate{500, 9}: Rock,
			Coordinate{499, 9}: Rock,
			Coordinate{498, 9}: Rock,
			Coordinate{497, 9}: Rock,
			Coordinate{496, 9}: Rock,
			Coordinate{495, 9}: Rock,
			Coordinate{494, 9}: Rock,
		}
		assert.Equal(t, expected, result)
	})

	t.Run("should return expected Abyss", func(t *testing.T) {
		assert.Equal(t, 9, limits.StartOfVoid)
	})

}
