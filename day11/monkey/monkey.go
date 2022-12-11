package monkey

import (
	"github.com/samber/lo"
)

type Monkey struct {
	Items           []int
	Operation       Operation
	TestModulus     int
	TargetTrue      int
	TargetFalse     int
	InspectionCount int
}

type MonkeyList []Monkey

type WorryLevelManagementStrategy func(int) int

func DivideStrategy(divisor int) WorryLevelManagementStrategy {
	return func(i int) int {
		return i / divisor
	}
}

func ModulusStrategy(monkeys MonkeyList) WorryLevelManagementStrategy {
	greatestCommonDivider := lo.Reduce(monkeys, func(agg int, item Monkey, _ int) int {
		return item.TestModulus * agg
	}, 1)

	return func(i int) int {
		return i % greatestCommonDivider
	}
}

func (l MonkeyList) NextRound(strategy WorryLevelManagementStrategy) MonkeyList {
	for idx := range l {
		monkey := &l[idx]
		var newMonkeyItems = []int{}
		for itemIdx := range monkey.Items {
			item := monkey.Items[itemIdx]

			newWorryLevel := monkey.Operation.Exec(item)
			newWorryLevelManaged := strategy(newWorryLevel)

			targetMonkeyIndex := lo.Ternary(newWorryLevelManaged%monkey.TestModulus == 0, monkey.TargetTrue, monkey.TargetFalse)

			targetMonkey := &l[targetMonkeyIndex]
			targetMonkey.Items = append(targetMonkey.Items, newWorryLevelManaged)
			monkey.InspectionCount++
		}
		monkey.Items = newMonkeyItems
	}
	return l
}