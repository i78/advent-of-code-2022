package cpu

type Cpu struct {
	X int
}

type SimulationResult map[int]Cpu

func (c Cpu) Simulate(program Program) SimulationResult {
	result := make(map[int]Cpu)
	cycle := 0
	currentCpu := c

	for _, opcode := range program {
		steps := opcode.Simulate(currentCpu)

		for _, s := range steps {
			currentCpu, result[cycle] = s, s
			cycle++
		}
	}

	return result
}

func NewCpu() Cpu {
	return Cpu{X: 1}
}

func (s *SimulationResult) CpuStateBeforeCycle(cycle int) Cpu {
	return (*s)[cycle-2]
}

func (s *SimulationResult) CpuStateAfterCycle(cycle int) Cpu {
	return (*s)[cycle-1]
}
