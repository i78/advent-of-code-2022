package treeGrid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsVisibleFromOutside(t *testing.T) {
	fakeGrid := Grid{
		{Height(3), Height(0), Height(3), Height(7), Height(3)},
		{Height(2), Height(5), Height(5), Height(1), Height(2)},
		{Height(6), Height(5), Height(3), Height(3), Height(2)},
		{Height(3), Height(3), Height(5), Height(4), Height(9)},
		{Height(3), Height(5), Height(3), Height(9), Height(0)},
	}

	testcases := []struct {
		x       int
		y       int
		visible bool
	}{
		{
			x:       1,
			y:       1,
			visible: true,
		}, {
			x:       2,
			y:       1,
			visible: true,
		}, {
			x:       3,
			y:       1,
			visible: false,
		}, {
			x:       2,
			y:       2,
			visible: false,
		}, {
			x:       4,
			y:       3,
			visible: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %t for %d,%d", testcase.visible, testcase.x, testcase.y), func(t *testing.T) {
			assert.Equal(t, testcase.visible, fakeGrid.IsVisibleFromOutside(testcase.x, testcase.y))
		})
	}

	t.Run("Should have expected total visible", func(t *testing.T) {
		visible := 0
		for y, row := range fakeGrid {
			for x := range row {
				if fakeGrid.IsVisibleFromOutside(x, y) {
					//fmt.Println(x, y)
					visible++
				}
			}
		}

		assert.Equal(t, 21, visible)
	})

}

func TestScenicScore(t *testing.T) {
	fakeGrid := Grid{
		{Height(3), Height(0), Height(3), Height(7), Height(3)},
		{Height(2), Height(5), Height(5), Height(1), Height(2)},
		{Height(6), Height(5), Height(3), Height(3), Height(2)},
		{Height(3), Height(3), Height(5), Height(4), Height(9)},
		{Height(3), Height(5), Height(3), Height(9), Height(0)},
	}

	testcases := []struct {
		x     int
		y     int
		score int
	}{
		{
			x:     2,
			y:     1,
			score: 4,
		}, {
			x:     2,
			y:     3,
			score: 8,
		},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("Should return %d for %d,%d", testcase.score, testcase.x, testcase.y), func(t *testing.T) {
			assert.Equal(t, testcase.score, fakeGrid.ScenicScore(testcase.x, testcase.y))
		})
	}

}
