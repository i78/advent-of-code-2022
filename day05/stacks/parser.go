package stacks

import (
	"strconv"
	"strings"
)

const LegendIndicator = " 1"
const TelegramSeparator = "\n"
const StackItemSeparator = '['
const ProgramLexSeparator = " "

const (
	ParserModeStacks = iota
	ParserModeSkipSectionSeparator
	ParserModeProgram
)

func NewStacks(raw string) (stacks Stacks, program Program) {
	mode := ParserModeStacks
	stacks = make(Stacks, 0, 10)
	for _, telegram := range strings.Split(raw, TelegramSeparator) {
		switch mode {
		case ParserModeStacks:
			if strings.HasPrefix(telegram, LegendIndicator) {
				mode = ParserModeSkipSectionSeparator
				continue
			}

			for at, offset := 0, 0; offset < len(telegram); at = strings.IndexByte(telegram[offset:], StackItemSeparator) {
				if at == -1 {
					break
				}
				offset += at
				if item := telegram[offset+1]; item != ' ' {
					targetStack := offset / 4
					stacks.PushBottom(targetStack, item)
				}
				offset++
			}

		case ParserModeSkipSectionSeparator:
			mode = ParserModeProgram

		case ParserModeProgram:
			lex := strings.Split(telegram, ProgramLexSeparator)
			program = append(program, ProgramStepMove{
				relaxedParseInt(lex[1]),
				relaxedParseInt(lex[3]),
				relaxedParseInt(lex[5])})
		}
	}
	return
}

func relaxedParseInt(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}
