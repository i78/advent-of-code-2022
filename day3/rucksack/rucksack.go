package rucksack

import (
	"errors"
	"github.com/samber/lo"
	"strings"
)

const TelegramSeparator = "\n"

type Rucksack struct {
	Compartments []Compartment
}

type Rucksacks []Rucksack
type Item = byte
type Compartment = []Item
type Items = []Item

func NewRucksackList(raw string) (result Rucksacks) {
	for _, tokens := range strings.Split(raw, TelegramSeparator) {
		itemsInRucksack := []byte(tokens)
		compartments := lo.Chunk[byte](itemsInRucksack, len(itemsInRucksack)/2)

		result = append(result, Rucksack{
			Compartments: compartments,
		})
	}
	return
}

func (rucksacks *Rucksacks) ToGroups(chunkSize int) (chunks []Rucksacks) {
	chunks = make([]Rucksacks, 0, len(*rucksacks)/chunkSize)

	for pos := 0; pos < len(*rucksacks); pos += chunkSize {
		chunks = append(chunks, (*rucksacks)[pos:pos+3])
	}

	return
}

func Priority(item Item) (priority int, err error) {
	const UppercasePrioStart, LowercasePrioStart = 27, 1
	switch {
	case item >= 'A' && item <= 'Z':
		priority = UppercasePrioStart + int(item-'A')
	case item >= 'a' && item <= 'z':
		priority = LowercasePrioStart + int(item-'a')
	default:
		err = errors.New("unknown item type")
	}
	return
}

func (r Rucksack) CommonItem() Item {
	return lo.Intersect[byte](r.Compartments[0], r.Compartments[1])[0]
}

func (rucksacks *Rucksacks) CommonItems() (items Items) {
	for _, rucksack := range *rucksacks {
		items = append(items, rucksack.CommonItem())
	}
	return
}

func TotalScore(items Items) (score int) {
	return lo.SumBy[Item](items, func(item Item) int {
		if s, err := Priority(item); err == nil {
			return s
		} else {
			panic("score error")
		}
	})
}

func (rucksacks *Rucksacks) CommonItemInGroup() Item {
	bothCompartmentsEach := lo.Map[Rucksack, Compartment](*rucksacks, func(r Rucksack, _ int) Compartment {
		return append(r.Compartments[0], r.Compartments[1]...)
	})

	common := lo.Uniq[byte](
		lo.Intersect[byte](bothCompartmentsEach[2],
			lo.Intersect[byte](bothCompartmentsEach[0], bothCompartmentsEach[1])))

	if len(common) > 1 {
		panic(">1 common")
	}

	return common[0]
}

func AllBadgeItemsFrom(groups []Rucksacks) (result Items) {
	for _, g := range groups {
		c := g.CommonItemInGroup()
		result = append(result, c)
	}
	return
}
