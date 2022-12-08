package treeGrid

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const ItemSeparator = ""

func NewGrid(raw string) (result Grid) {
	for _, line := range strings.Split(raw, LineSeparator) {
		tokens := strings.Split(line, ItemSeparator)
		result = append(result, lo.Map[string, Height](tokens, func(token string, _ int) Height {
			b, _ := strconv.Atoi(token)
			return Height(b)
		}))
	}
	return
}
