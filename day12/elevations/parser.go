package elevations

import (
	"github.com/samber/lo"
	"strings"
)

const LineSeparator = "\n"

func NewElevationsMap(raw string) (result ElevationMap, start Coordinate, destination Coordinate) {
	lines := strings.Split(raw, LineSeparator)

	for lineIdx, line := range lines {
		ch := strings.Split(line, "")
		bytes := lo.Map(ch, func(item string, idxx int) byte {
			at := item[0]
			if at == 'S' {
				at = 'a'
				start = Coordinate{X: idxx, Y: lineIdx}
			} else if at == 'E' {
				at = 'z'
				destination = Coordinate{X: idxx, Y: lineIdx}
			}
			return at
		})
		result = append(result, bytes)
	}

	return
}
