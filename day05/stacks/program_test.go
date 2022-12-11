package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgram(t *testing.T) {

	setup := func() (Stacks, Program) {
		return Stacks{
				Stack{'N', 'Z'},
				Stack{'D', 'C', 'M'},
				Stack{'P'},
			},
			Program{
				ProgramStepMove{1, 2, 1},
				ProgramStepMove{3, 1, 3},
				ProgramStepMove{2, 2, 1},
				ProgramStepMove{1, 1, 2},
			}
	}

	t.Run("Given Program should return expected result for CrateMover9000", func(t *testing.T) {
		stacks, program := setup()

		stacks = program.Run(stacks, CrateMover9000, nil)

		assert.Equal(t, Stacks{
			Stack{'C'},
			Stack{'M'},
			Stack{'Z', 'N', 'D', 'P'},
		}, stacks)
	})

	t.Run("Given Program should return expected result for CrateMover9001", func(t *testing.T) {
		stacks, program := setup()

		stacks = program.Run(stacks, CrateMover9001, nil)

		assert.Equal(t, Stacks{
			Stack{'M'},
			Stack{'C'},
			Stack{'D', 'N', 'Z', 'P'},
		}, stacks)
	})
}
