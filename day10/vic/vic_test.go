package vic

import (
	"dreese.de/aoc22/day10/cpu"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestVic(t *testing.T) {

	t.Run("with test program", func(t *testing.T) {
		var testProgram cpu.Program
		if content, err := os.ReadFile("../testprogram.txt"); err != nil {
			t.Fatal("Cannot load test program!")
		} else {
			testProgram = cpu.NewProgram(string(content))
		}

		t.Run("should return expected image", func(t *testing.T) {
			result := cpu.NewCpu().Simulate(testProgram)

			const crtWidth = 40
			const expectedImage = "\n0 \t ▆▆  ▆▆  ▆▆  ▆▆  ▆▆  ▆▆  ▆▆  ▆▆  ▆▆  ▆▆   39 \n40 \t ▆▆▆   ▆▆▆   ▆▆▆   ▆▆▆   ▆▆▆   ▆▆▆   ▆▆▆  79 \n80 \t ▆▆▆▆    ▆▆▆▆    ▆▆▆▆    ▆▆▆▆    ▆▆▆▆     119 \n120 \t ▆▆▆▆▆     ▆▆▆▆▆     ▆▆▆▆▆     ▆▆▆▆▆      159 \n160 \t ▆▆▆▆▆▆      ▆▆▆▆▆▆      ▆▆▆▆▆▆      ▆▆▆▆ 199 \n200 \t ▆▆▆▆▆▆▆       ▆▆▆▆▆▆▆       ▆▆▆▆▆▆▆     "

			screen := SimulateVic(result, crtWidth)

			assert.Equal(t, expectedImage, screen)
		})
	})
}
