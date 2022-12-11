package monkey

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testData = "Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1"

func TestParseMonkeys(t *testing.T) {
	t.Run("should return monkeys as expected", func(t *testing.T) {
		result := NewMonkeyList(testData)

		expectedResult := MonkeyList{
			Monkey{
				Items:       []int{79, 98},
				Operation:   Mul{Operand: 19},
				TestModulus: 23,
				TargetTrue:  2,
				TargetFalse: 3,
			}, Monkey{
				Items:       []int{54, 65, 75, 74},
				Operation:   Add{Operand: 6},
				TestModulus: 19,
				TargetTrue:  2,
				TargetFalse: 0,
			}, Monkey{
				Items:       []int{79, 60, 97},
				Operation:   Pow2{Operand: 0},
				TestModulus: 13,
				TargetTrue:  1,
				TargetFalse: 3,
			}, Monkey{
				Items:       []int{74},
				Operation:   Add{Operand: 3},
				TestModulus: 17,
				TargetTrue:  0,
				TargetFalse: 1,
			},
		}

		assert.Equal(t, expectedResult, result)
	})
}

func TestParser(t *testing.T) {
	testcases := []struct {
		input    string
		expected Operation
	}{
		{input: "new = old * 19", expected: Mul{19}},
		{input: "new = old + 6", expected: Add{6}},
		{input: "new = old * 19", expected: Mul{19}},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %v for (%s)", testcase.expected, testcase.input), func(t *testing.T) {
			result := parse(testcase.input)
			assert.Equal(t, testcase.expected, result)
		})
	}
}