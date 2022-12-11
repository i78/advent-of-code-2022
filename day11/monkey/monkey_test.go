package monkey

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMonkeyRound(t *testing.T) {

	fakeMonkeys := func() MonkeyList { return MonkeyList{
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
	}}

	t.Run("should have expected state after one round", func(t *testing.T) {
		monkeys := fakeMonkeys()

		strategy := DivideStrategy(3)

		subject := monkeys.NextRound(strategy)

		assert.Equal(t,  []int {20,23,27,26}, subject[0].Items)
		assert.Equal(t,  []int {2080, 25 , 167, 207 , 401, 1046}, subject[1].Items)
		assert.Equal(t,  []int {}, subject[2].Items)
		assert.Equal(t,  []int {}, subject[3].Items)
	})

	t.Run("should have expected state after two rounds", func(t *testing.T) {
		monkeys := fakeMonkeys()
		strategy := DivideStrategy(3)

		subject := monkeys.NextRound(strategy)
		subject = monkeys.NextRound(strategy)

		assert.Equal(t,  []int {695, 10, 71, 135, 350}, subject[0].Items)
		assert.Equal(t,  []int {43, 49, 58, 55, 362}, subject[1].Items)
		assert.Equal(t,  []int {}, subject[2].Items)
		assert.Equal(t,  []int {}, subject[3].Items)
	})

	t.Run("After 20 rounds", func(t *testing.T) {
		monkeys := fakeMonkeys()

		subject := monkeys

		for i := 0 ; i < 20 ; i++ {
			subject = monkeys.NextRound(DivideStrategy(3))
		}

		t.Run("monkeys should have expected items", func(t *testing.T) {
			assert.Equal(t,  []int {10, 12, 14, 26, 34}, subject[0].Items)
			assert.Equal(t,  []int {245, 93, 53, 199, 115}, subject[1].Items)
			assert.Equal(t,  []int {}, subject[2].Items)
			assert.Equal(t,  []int {}, subject[3].Items)
		})

		t.Run("monkeys should have inspected the expected number of items", func(t *testing.T) {
			assert.Equal(t,  101, subject[0].InspectionCount)
			assert.Equal(t,  95, subject[1].InspectionCount)
			assert.Equal(t,  7, subject[2].InspectionCount)
			assert.Equal(t,  105, subject[3].InspectionCount)
		})
	})

	t.Run("After 10000 rounds", func(t *testing.T) {
		monkeys := fakeMonkeys()

		for i := 0 ; i < 10000 ; i++ {
			monkeys = monkeys.NextRound(ModulusStrategy(monkeys))
		}

		t.Run("monkeys should have inspected the expected number of items", func(t *testing.T) {
			assert.Equal(t,  52166, monkeys[0].InspectionCount)
			assert.Equal(t,  47830, monkeys[1].InspectionCount)
			assert.Equal(t,  1938, monkeys[2].InspectionCount)
			assert.Equal(t,  52013, monkeys[3].InspectionCount)
		})
	})

}