package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinct2(t *testing.T) {
	testcases := []struct {
		input          string
		width          int
		expectedResult int
	}{
		{
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			width:          4,
			expectedResult: 7,
		}, {
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			width:          4,
			expectedResult: 5,
		}, {
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			width:          4,
			expectedResult: 6,
		}, {
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			width:          4,
			expectedResult: 10,
		}, {
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			width:          4,
			expectedResult: 11,
		}, {
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			width:          14,
			expectedResult: 19,
		}, {
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			width:          14,
			expectedResult: 23,
		}, {
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			width:          14,
			expectedResult: 23,
		}, {
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			width:          14,
			expectedResult: 29,
		}, {
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			width:          14,
			expectedResult: 26,
		}, {
			input:          "a",
			width:          14,
			expectedResult: -1,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %d for s=%s w=%d", testcase.expectedResult, testcase.input, testcase.width), func(t *testing.T) {
			assert.Equal(t, testcase.expectedResult, FirstDistinct(testcase.width, []byte(testcase.input)))
		})
	}
}
