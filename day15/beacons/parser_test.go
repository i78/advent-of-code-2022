package beacons

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15\nSensor at x=9, y=16: closest beacon is at x=10, y=16\nSensor at x=13, y=2: closest beacon is at x=15, y=3\nSensor at x=12, y=14: closest beacon is at x=10, y=16\nSensor at x=10, y=20: closest beacon is at x=10, y=16\nSensor at x=14, y=17: closest beacon is at x=10, y=16\nSensor at x=8, y=7: closest beacon is at x=2, y=10\nSensor at x=2, y=0: closest beacon is at x=2, y=10\nSensor at x=0, y=11: closest beacon is at x=2, y=10\nSensor at x=20, y=14: closest beacon is at x=25, y=17\nSensor at x=17, y=20: closest beacon is at x=21, y=22\nSensor at x=16, y=7: closest beacon is at x=15, y=3\nSensor at x=14, y=3: closest beacon is at x=15, y=3\nSensor at x=20, y=1: closest beacon is at x=15, y=3"

func TestParser(t *testing.T) {
	parser := participle.MustBuild[ReadingList](
		participle.Lexer(SensorReadingListLexer),
	)

	ini, err := parser.ParseString("", testdata)
	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
	assert.NoError(t, err)
}
