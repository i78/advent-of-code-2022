package signals

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	testcases := []struct {
		signal   string
		expected bool
	}{
		{
			signal:   "[1,1,3,1,1]\n[1,1,5,1,1]",
			expected: true,
		}, {
			signal:   "[[1],[2,3,4]]\n[[1],4]",
			expected: true,
		}, {
			signal:   "[9]\n[[8,7,6]]",
			expected: false,
		}, {
			signal:   "[[4,4],4,4]\n[[4,4],4,4,4]",
			expected: true,
		}, {
			signal:   "[7,7,7,7]\n[7,7,7]",
			expected: false,
		}, {
			signal:   "[]\n[3]",
			expected: true,
		}, {
			signal:   "[[[]]]\n[[]]",
			expected: false,
		}, {
			signal:   "[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]",
			expected: false,
		}, {
			signal:   "[[[]]]\n[1,1,3,1,1]",
			expected: true,
		}, {
			signal:   "[[1],4]\n[[1],[2,3,4]]",
			expected: false,
		},
	}
	for _, test := range testcases {
		t.Run(fmt.Sprintf("should return %t for %s", test.expected, test.signal), func(t *testing.T) {
			data := NewSignalsList(test.signal)
			pair := data[0]

			assert.Equal(t, test.expected, pair.InOrder())
		})
	}
}
