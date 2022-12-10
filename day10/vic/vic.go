package vic

import (
	"dreese.de/aoc22/day10/cpu"
	"fmt"
	"github.com/samber/lo"
	"math"
	"strings"
)

func SimulateVic(simulationResult cpu.SimulationResult, crtWidth int) string {
	var builder strings.Builder
	for i := 0; i < len(simulationResult); i++ {
		column := i % crtWidth

		if column == 0 {
			if i != 0 {
				builder.WriteString(fmt.Sprintf(" %d ", i-1))
			}
			builder.WriteString(fmt.Sprintf("\n%d \t ", i))
		}

		xRegisterAfterCycle := simulationResult.CpuStateAfterCycle(i).X
		builder.WriteString(lo.Ternary(math.Abs(float64(xRegisterAfterCycle-column)) < 2, "▆", " "))
	}
	return builder.String()
}
