package stacks

import (
	"github.com/samber/lo"
)

type Program []ProgramStepMove

type ProgramStepMove struct {
	Quantity int
	From     int
	To       int
}

type Machine byte

const (
	CrateMover9000 Machine = iota
	CrateMover9001
)

func (p Program) Run(stacks Stacks, machine Machine, after ...func(stacks *Stacks)) Stacks {
	for _, step := range p {
		switch machine {
		case CrateMover9000:
			stacks = step.ExecSingle(stacks)
		case CrateMover9001:
			stacks = step.ExecMultiple(stacks)
		default:
			panic("unknown machine!")
		}

		for _, fnAfter := range after {
			fnAfter(&stacks)
		}
	}
	return stacks
}

func (m ProgramStepMove) ExecSingle(stacks Stacks) Stacks {
	for i := 1; i <= m.Quantity; i++ {
		if item, err := stacks.Pop(m.From - 1); err == nil {
			stacks.PushTop(m.To-1, item)
		}
	}
	return stacks
}

func (m ProgramStepMove) ExecMultiple(stacks Stacks) Stacks {
	items := make([]uint8, 0, 100)

	for i := 1; i <= m.Quantity; i++ {
		if item, err := stacks.Pop(m.From - 1); err == nil {
			items = append(items, item)
		}
	}

	for _, item := range lo.Reverse[uint8](items) {
		stacks.PushTop(m.To-1, item)
	}

	return stacks
}
