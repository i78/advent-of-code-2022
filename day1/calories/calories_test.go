package calories

import "testing"
import assert "github.com/stretchr/testify/assert"

const testdata = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"

func TestWhenTestDataParsed(t *testing.T) {

	t.Run("Should return result with expected cardinality of 5 (elves)", func(t *testing.T) {
		elvesCaloriesRecords := NewCaloryAmounts(testdata)

		const expectedNumberOfElves = 5
		assert.Len(t, elvesCaloriesRecords, expectedNumberOfElves)
	})

	t.Run("Should contain expected calory count for elves", func(t *testing.T) {
		expectedCaloryValues := []int{6000, 4000, 11000, 24000, 10000}

		elvesCaloriesRecords := NewCaloryAmounts(testdata)

		assert.Equal(t, expectedCaloryValues, elvesCaloriesRecords)
	})
}

func TestWhenAskedForMaxCaloryValue(t *testing.T) {
	t.Run("Should return max value from array", func(t *testing.T) {
		fakeCaloryValues := []CaloryAmount{6000, 4000, 11000, 24000, 10000}

		maxCalories := FindMaxCaloryValueCarriedByElv(fakeCaloryValues)

		const expectedMaxValue = 24000
		assert.Equal(t, maxCalories, expectedMaxValue)
	})
}

func TestCaloriesCarriedByTop3Elves(t *testing.T) {
	t.Run("Should return max value from array", func(t *testing.T) {
		fakeCaloryValues := []int{6000, 4000, 11000, 24000, 10000}
		const expectedMaxValue = 45000

		totalTop3Calories := CaloriesCarriedByTop3Elves(fakeCaloryValues)

		assert.Equal(t, expectedMaxValue, totalTop3Calories)
	})
}
