package stacks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRearrange(t *testing.T) {
	t.Run("Should return stack in expected State", func(t *testing.T) {

		testcases := []struct {
			stacks         Stacks
			operation      ProgramStepMove
			expectedStacks Stacks
		}{
			{
				stacks: Stacks{
					Stack{'N', 'Z'},
					Stack{'D', 'C', 'M'},
					Stack{'P'},
				},
				operation: ProgramStepMove{1, 2, 1},
				expectedStacks: Stacks{
					Stack{'D', 'N', 'Z'},
					Stack{'C', 'M'},
					Stack{'P'},
				},
			}, {
				stacks: Stacks{
					Stack{'D', 'N', 'Z'},
					Stack{'C', 'M'},
					Stack{'P'},
				},
				operation: ProgramStepMove{3, 1, 3},
				expectedStacks: Stacks{
					Stack{},
					Stack{'C', 'M'},
					Stack{'Z', 'N', 'D', 'P'},
				},
			}, {
				stacks: Stacks{
					Stack{},
					Stack{'C', 'M'},
					Stack{'Z', 'N', 'D', 'P'},
				},
				operation: ProgramStepMove{2, 2, 1},
				expectedStacks: Stacks{
					Stack{'M', 'C'},
					Stack{},
					Stack{'Z', 'N', 'D', 'P'},
				},
			},
		}

		for idx, testcase := range testcases {
			t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
				result := testcase.operation.ExecSingle(testcase.stacks)

				assert.Equal(t, testcase.expectedStacks, result)
				fmt.Println(result)
			})
		}

	})
}
