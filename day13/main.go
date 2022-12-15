package main

import (
	"dreese.de/aoc22/day13/signals"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"log"
	"os"
)

func main() {
	if content, err := os.ReadFile("input.txt"); err == nil {
		fmt.Printf("Part 1: %d\nPart 2: %d", SolveA(content), SolveB(content))
	} else {
		log.Fatal(err)
	}
}

func SolveA(content []byte) int {
	pairs := signals.NewSignalsList(string(content))

	return lo.Sum(lo.Map(pairs, func(pair signals.Pair, idx int) int {
		return lo.Ternary(pair.InOrder(), idx, 0)
	}))
}

func SolveB(content []byte) int {
	dividerPackets := signals.NewSignalsList("[[2]]\n[[6]]")
	pairs := append(signals.NewSignalsList(string(content)), dividerPackets[0])

	packets := lo.Flatten(lo.Map(pairs, func(pair signals.Pair, _ int) []*signals.Node {
		return []*signals.Node{pair.Left, pair.Right}
	}))

	slices.SortFunc(packets, func(a *signals.Node, b *signals.Node) bool {
		return (&signals.Pair{a, b}).InOrder()
	})

	return lo.Reduce(packets, func(agg int, it *signals.Node, idx int) int {
		return lo.Ternary(it == dividerPackets[0].Left || it == dividerPackets[0].Right,
			agg*(idx+1), agg)
	}, 1)
}
