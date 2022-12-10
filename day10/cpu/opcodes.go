package cpu

type Program []Opcode

type Opcode interface {
	Simulate(cpu Cpu) []Cpu
}

type Noop struct {
}

func (n Noop) Simulate(cpu Cpu) []Cpu {
	return []Cpu{cpu}
}

type Addx struct {
	N int
}

func (n Addx) Simulate(cpu Cpu) []Cpu {
	return []Cpu{cpu, {X: cpu.X + n.N}}
}
