package signals

import (
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const PairSeparator = "\n\n"

func ToTree(parent *Node, signal []byte) *Node {
	result := &Node{
		parent:   parent,
		children: []*Node{},
	}

	jsonparser.ArrayEach(signal, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if dataType == jsonparser.Array {
			result.children = append(result.children, ToTree(result, value))
		} else {
			i, _ := strconv.Atoi(string(value))
			result.children = append(result.children, &Node{
				parent:   result,
				value:    &i,
				children: nil,
			})
		}
	})
	return result
}

func NewSignalsList(raw string) (result []Pair) {
	for _, pair := range strings.Split(raw, PairSeparator) {
		lines := strings.Split(pair, LineSeparator)
		result = append(result, Pair{
			Left:  ToTree(nil, []byte(lines[0])),
			Right: ToTree(nil, []byte(lines[1])),
		})
	}
	return
}
