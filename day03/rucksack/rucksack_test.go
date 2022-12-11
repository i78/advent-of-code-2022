package rucksack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"

func TestWhenTestDataParsed(t *testing.T) {

	t.Run("Should return result with expected count of 6 (rucksacks)", func(t *testing.T) {
		subject := NewRucksackList(testdata)

		const expectedNumberOfRounds = 6
		assert.Len(t, subject, expectedNumberOfRounds)
	})

	t.Run("Should have expected compartment contents", func(t *testing.T) {
		testcases := []struct {
			index        int
			compartments []Compartment
		}{
			{index: 0, compartments: []Compartment{
				{'v', 'J', 'r', 'w', 'p', 'W', 't', 'w', 'J', 'g', 'W', 'r'},
				{'h', 'c', 's', 'F', 'M', 'M', 'f', 'F', 'F', 'h', 'F', 'p'},
			}},
			{index: 1, compartments: []Compartment{
				{'j', 'q', 'H', 'R', 'N', 'q', 'R', 'j', 'q', 'z', 'j', 'G', 'D', 'L', 'G', 'L'},
				{'r', 's', 'F', 'M', 'f', 'F', 'Z', 'S', 'r', 'L', 'r', 'F', 'Z', 's', 'S', 'L'},
			}}, {index: 2, compartments: []Compartment{
				{'P', 'm', 'm', 'd', 'z', 'q', 'P', 'r', 'V'},
				{'v', 'P', 'w', 'w', 'T', 'W', 'B', 'w', 'g'},
			}},
		}

		subject := NewRucksackList(testdata)

		for _, testcase := range testcases {
			t.Run(fmt.Sprintf("Should return expected result at idx %d", testcase.index), func(t *testing.T) {
				assert.Equal(t, testcase.compartments, subject[testcase.index].Compartments)
			})
		}

	})
}

func TestItemPriorities(t *testing.T) {
	testcases := []struct {
		item     Item
		priority int
	}{
		{item: 'p', priority: 16},
		{priority: 38, item: 'L'},
		{priority: 42, item: 'P'},
		{priority: 22, item: 'v'},
		{priority: 20, item: 't'},
		{priority: 19, item: 's'},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return prio=%d for item=%c", testcase.priority, testcase.item), func(t *testing.T) {
			subject, err := Priority(testcase.item)
			assert.NoErrorf(t, err, "...")
			assert.Equal(t, testcase.priority, subject)
		})
	}
}

func TestCommonItem(t *testing.T) {
	testcases := []struct {
		commonItem Item
		rucksack   Rucksack
	}{
		{commonItem: 'p', rucksack: Rucksack{Compartments: []Compartment{
			{'v', 'J', 'r', 'w', 'p', 'W', 't', 'w', 'J', 'g', 'W', 'r'},
			{'h', 'c', 's', 'F', 'M', 'M', 'f', 'F', 'F', 'h', 'F', 'p'},
		}}},
		{commonItem: 'L', rucksack: Rucksack{Compartments: []Compartment{
			{'j', 'q', 'H', 'R', 'N', 'q', 'R', 'j', 'q', 'z', 'j', 'G', 'D', 'L', 'G', 'L'},
			{'r', 's', 'F', 'M', 'f', 'F', 'Z', 'S', 'r', 'L', 'r', 'F', 'Z', 's', 'S', 'L'},
		}}}, {commonItem: 'P', rucksack: Rucksack{Compartments: []Compartment{
			{'P', 'm', 'm', 'd', 'z', 'q', 'P', 'r', 'V'},
			{'v', 'P', 'w', 'w', 'T', 'W', 'B', 'w', 'g'},
		}}},
	}

	for idx, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return commonItem=%c for rucksack=%c", testcase.commonItem, idx), func(t *testing.T) {
			subject := testcase.rucksack.CommonItem()
			assert.Equal(t, testcase.commonItem, subject)
		})
	}

	t.Run("Should return expected List of common items", func(t *testing.T) {
		expectedCommonItems := Items{'p', 'L', 'P'}
		allRucksacks := Rucksacks{testcases[0].rucksack, testcases[1].rucksack, testcases[2].rucksack}
		assert.Equal(t, expectedCommonItems, allRucksacks.CommonItems())
	})
}

func TestCommonItemScore(t *testing.T) {

	t.Run("should return score 157 for given common items", func(t *testing.T) {
		items := Items{'p', 'L', 'P', 'v', 't', 's'}

		score := TotalScore(items)
		assert.Equal(t, 157, score)
	})
}

func TestFindBadges(t *testing.T) {
	input := NewRucksackList("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw")

	t.Run("Should return two groups of 3 elves", func(t *testing.T) {
		groups := input.ToGroups(3)
		assert.Len(t, groups, 2)
		assert.Len(t, groups[0], 3)
		assert.Len(t, groups[1], 3)
	})

	t.Run("Should find common item 'r' in first group", func(t *testing.T) {
		groups := input.ToGroups(3)
		commonItem := groups[0].CommonItemInGroup()
		assert.Equal(t, Item('r'), commonItem)
	})

	t.Run("Should find common item 'Z' second group", func(t *testing.T) {
		groups := input.ToGroups(3)
		commonItem := groups[1].CommonItemInGroup()
		assert.Equal(t, Item('Z'), commonItem)
	})

	t.Run("Should return the expected sum priority of 70", func(t *testing.T) {
		groups := input.ToGroups(3)

		totalItems := AllBadgeItemsFrom(groups)
		score := TotalScore(totalItems)
		assert.Equal(t, Items{'r', 'Z'}, totalItems)
		assert.Equal(t, 70, score)
	})
}
