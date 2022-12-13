package elevations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testData = "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"

func TestParseMap(t *testing.T) {
	t.Run("should return map as expected", func(t *testing.T) {
		result, start, destination := NewElevationsMap(testData)

		expectedResult := ElevationMap{
			[]byte{0x61, 0x61, 0x62, 0x71, 0x70, 0x6f, 0x6e, 0x6d},
			[]byte{0x61, 0x62, 0x63, 0x72, 0x79, 0x78, 0x78, 0x6c},
			[]byte{0x61, 0x63, 0x63, 0x73, 0x7a, 0x7a, 0x78, 0x6b},
			[]byte{0x61, 0x63, 0x63, 0x74, 0x75, 0x76, 0x77, 0x6a},
			[]byte{0x61, 0x62, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69}}

		assert.Equal(t, expectedResult, result)
		assert.Equal(t, Coordinate{0, 0}, start)
		assert.Equal(t, Coordinate{5, 2}, destination)
	})

}
