package cave

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDropSand(t *testing.T) {

	t.Run("should drop to 500,8 when deployed at 500,0", func(t *testing.T) {
		cave, limits := NewCave(testdata)
		endPosition, droppedIntoAbyss := cave.Drop(limits)

		assert.Equal(t, Coordinate{500, 8}, endPosition)
		assert.False(t, droppedIntoAbyss)
	})

	t.Run("should drop to 499,8 when deployed 2nd at 500,0", func(t *testing.T) {
		cave, limits := NewCave(testdata)
		cave.Drop(limits)
		endPosition, droppedIntoAbyss := cave.Drop(limits)

		assert.Equal(t, Coordinate{499, 8}, endPosition)
		assert.False(t, droppedIntoAbyss)
	})

	t.Run("should report drop to void after 24 drops", func(t *testing.T) {
		cave, limits := NewCave(testdata)
		for i := 0; i < 24; i++ {
			endPosition, droppedIntoAbyss := cave.Drop(limits)
			fmt.Printf("%d %v %t\n", i, endPosition, droppedIntoAbyss)

			if i == 24 {
				assert.True(t, droppedIntoAbyss)
			}
		}
	})
}

func TestDropSandWithFloorHeightSet(t *testing.T) {
	t.Run("should tetris kill the cave after 93 grains", func(t *testing.T) {
		cave, limits := NewCave(testdata)

		limits = limits.ToFloor()

		for i := 0; i <= 92; i++ {
			endPosition, failedToDescentFrom500 := cave.Drop(limits)
			fmt.Printf("%d %v %t\n", i, endPosition, failedToDescentFrom500)

			if i == 92 {
				assert.True(t, failedToDescentFrom500)
			}
		}

	})
}
