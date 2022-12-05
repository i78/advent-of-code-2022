package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testdata = "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2"

func TestWhenTestDataParsed(t *testing.T) {

	t.Run("Should return result with expected count of 3 stacks", func(t *testing.T) {
		stacks, _ := NewStacks(testdata)

		const expectedSize = 3
		assert.Len(t, stacks, expectedSize)
	})

	t.Run("Should return expected Stack", func(t *testing.T) {
		stacks, _ := NewStacks(testdata)

		expectedStack := Stacks{
			Stack{'N', 'Z'},
			Stack{'D', 'C', 'M'},
			Stack{'P'},
		}

		assert.Equal(t, expectedStack, stacks)
	})

	t.Run("Should return result with expected count of 4 program steps", func(t *testing.T) {
		_, program := NewStacks(testdata)

		const expectedSize = 4
		assert.Len(t, program, expectedSize)
	})

	t.Run("Should return expected program", func(t *testing.T) {
		_, program := NewStacks(testdata)

		expectedProgram := Program{
			ProgramStepMove{1, 2, 1},
			ProgramStepMove{3, 1, 3},
			ProgramStepMove{2, 2, 1},
			ProgramStepMove{1, 1, 2},
		}

		assert.Equal(t, expectedProgram, program)
	})
}
