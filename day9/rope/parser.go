package rope

import (
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const ItemSeparator = " "

func NewMovesList(raw string) (result Moves) {
	for _, line := range strings.Split(raw, LineSeparator) {
		tokens := strings.Split(line, ItemSeparator)
		steps, _ := strconv.Atoi(tokens[1])
		result = append(result, Move{
			Direction: Direction(tokens[0][0]),
			Steps:     steps,
		})
	}
	return
}
