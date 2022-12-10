package cpu

import (
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const ItemSeparator = " "

func NewProgram(raw string) (result Program) {
	for _, line := range strings.Split(raw, LineSeparator) {
		tokens := strings.Split(line, ItemSeparator)
		switch tokens[0] {
		case "noop":
			result = append(result, Noop{})
		case "addx":
			n, _ := strconv.Atoi(tokens[1])
			result = append(result, Addx{n})
		}
	}
	return
}
