package stacks

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
)

type StackItem byte
type Stack []byte
type Stacks []Stack

func (stacks *Stacks) PushBottom(idx int, item uint8) {
	stacks.ensureStackExists(idx)
	(*stacks)[idx] = append((*stacks)[idx], item)
}

func (stacks *Stacks) PushTop(idx int, item uint8) {
	stacks.ensureStackExists(idx)
	(*stacks)[idx] = append([]uint8{item}, (*stacks)[idx]...)
}

func (stacks *Stacks) ensureStackExists(idx int) *Stacks {
	if len(*stacks) < idx+1 {
		*stacks = (*stacks)[0 : idx+1]
	}
	if (*stacks)[idx] == nil {
		(*stacks)[idx] = make(Stack, 0, 10)
	}
	return stacks
}

func (stacks *Stacks) Pop(idx int) (item uint8, err error) {
	if item, err = stacks.PeekTop(idx); err == nil {
		(*stacks)[idx] = (*stacks)[idx][1:]
	}
	return
}

func (stacks *Stacks) PeekTop(idx int) (item uint8, err error) {
	if len(*stacks) < idx+1 {
		panic("Stack does not exist")
	}
	if len((*stacks)[idx]) == 0 {
		err = errors.New("stack underrun")
		return
	}
	item = (*stacks)[idx][0]
	return
}

func (stacks *Stacks) Print() {
	for stackId := range *stacks {
		top, _ := stacks.PeekTop(stackId)

		j := lo.Reduce[byte, string]((*stacks)[stackId], func(agg string, item byte, _ int) string {
			return agg + string(item)
		}, "")

		fmt.Printf("%60v %d %-1c\n", j, stackId, top)
	}
	fmt.Println("")
}
