package sections

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"

func TestWhenTestDataParsed(t *testing.T) {
	expectedResult := []SectionRangePair{
		{{2, 4}, {6, 8}},
		{{2, 3}, {4, 5}},
		{{5, 7}, {7, 9}},
		{{2, 8}, {3, 7}},
		{{6, 6}, {4, 6}},
		{{2, 6}, {4, 8}},
	}

	t.Run("Should return expected result", func(t *testing.T) {
		subject := NewSectionList(testdata)
		assert.Equal(t, expectedResult, subject)
	})
}

func TestFullOverlap(t *testing.T) {
	testCases := []struct {
		SectionRangePair
		fullyIncludes bool
		overlaps      bool
	}{
		{SectionRangePair: SectionRangePair{SectionRange{
			From: 1, To: 10}, SectionRange{From: 20, To: 30}},
			fullyIncludes: false, overlaps: false},
		{SectionRangePair: SectionRangePair{SectionRange{
			From: 1, To: 10}, SectionRange{From: 5, To: 9}},
			fullyIncludes: true, overlaps: true},
		{SectionRangePair: SectionRangePair{SectionRange{
			From: 1, To: 10}, SectionRange{From: 5, To: 15}},
			fullyIncludes: false, overlaps: true},
	}

	for idx, testcase := range testCases {
		t.Run(fmt.Sprintf("Should return %v/%v for case %d w", testcase.fullyIncludes, testcase.overlaps, idx), func(t *testing.T) {
			fullyIncludes, overlaps := testcase.SectionRangePair.Overlaps(FullyIncludes),
				testcase.SectionRangePair.Overlaps(Overlaps)

			assert.Equal(t, fullyIncludes, testcase.fullyIncludes)
			assert.Equal(t, overlaps, testcase.overlaps)
		})
	}

	t.Run("Should return 2 for example data and full overlap", func(t *testing.T) {
		subject := NewSectionList(testdata)

		numOfFullOverlaps := 0
		for _, it := range subject {
			if it.Overlaps(FullyIncludes) == true {
				numOfFullOverlaps++
			}
		}

		assert.Equal(t, 2, numOfFullOverlaps)
	})
}
