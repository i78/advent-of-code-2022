package cpu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSimulateCpu(t *testing.T) {

	t.Run("should have register X=1 right away", func(t *testing.T) {
		assert.Equal(t, 1, NewCpu().X)
	})

	t.Run("should simulate NOP program with 1 step and no register change", func(t *testing.T) {
		cpu := NewCpu()
		result := cpu.Simulate(Program{Noop{}})

		expected := SimulationResult{
			0: {
				X: 1,
			},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("should have expected cpu state in last simulation snapshot", func(t *testing.T) {
		cpu := NewCpu()
		program := Program{
			Noop{},
			Addx{3},
			Addx{-5},
		}

		result := cpu.Simulate(program)

		expectedLastState := Cpu{
			X: -1,
		}

		assert.Equal(t, expectedLastState, result[len(result)-1])
	})

	t.Run("with test program", func(t *testing.T) {
		var testProgram Program
		if content, err := os.ReadFile("../testprogram.txt"); err != nil {
			t.Fatal("Cannot load test program!")
		} else {
			testProgram = NewProgram(string(content))
		}

		t.Run("should return n>1 simulated steps", func(t *testing.T) {
			cpu := NewCpu()
			result := cpu.Simulate(testProgram)

			assert.Greater(t, len(result), 0)
		})

		expectedSignalStrengths := []struct {
			cycle    int
			expected int
		}{
			{20, 420},
			{60, 1140},
			{100, 1800},
			{140, 2940},
			{180, 2880},
			{220, 3960},
		}

		for _, testcase := range expectedSignalStrengths {
			t.Run(fmt.Sprintf("Should return signal strength %d at %d", testcase.expected, testcase.cycle), func(t *testing.T) {
				cpu := NewCpu()
				result := cpu.Simulate(testProgram)

				signalStrength := result.CpuStateBeforeCycle(testcase.cycle).X * testcase.cycle

				assert.Equal(t, testcase.expected, signalStrength)
			})
		}
	})
}
