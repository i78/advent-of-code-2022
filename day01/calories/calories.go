package calories

import (
	"github.com/samber/lo"
	"sort"
	"strconv"
	"strings"
)

const TelegramSeparator = "\n\n"
const ItemSeparator = "\n"

type CaloryAmount = int
type CaloryAmounts = []CaloryAmount

func NewCaloryAmounts(s string) (result []CaloryAmount) {
	for _, elv := range strings.Split(s, TelegramSeparator) {
		records := strings.Split(elv, ItemSeparator)

		totalCaloriesOfRecord := lo.SumBy[string, int](records, func(it string) int {
			if itemInt, err := strconv.ParseInt(it, 10, 32); err == nil {
				return int(itemInt)
			}
			return 0
		})

		result = append(result, totalCaloriesOfRecord)
	}
	return
}

func FindMaxCaloryValueCarriedByElv(c CaloryAmounts) CaloryAmount {
	return lo.Max[int](c)
}

func CaloriesCarriedByTop3Elves(groupAmounts []CaloryAmount) CaloryAmount {
	sorted := append([]int{}, groupAmounts...)
	sort.Ints(sorted)
	return lo.Sum[int](sorted[len(sorted)-3:])
}
