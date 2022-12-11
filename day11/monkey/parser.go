package monkey

import (
	"dreese.de/aoc22/day11/monkey/tokens"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

const MonkeySeparator = "\n\n"
const LineSeparator = "\n"
const TokenSeparator = ":"

func NewMonkeyList(raw string) (result MonkeyList) {
	for _, line := range strings.Split(raw, MonkeySeparator) {
		monkeys := strings.Split(line, MonkeySeparator)

		for _, m := range monkeys {
			var monkey Monkey
			monkeyLines := strings.Split(m, LineSeparator)

			for _, line := range monkeyLines {
				tokens := strings.Split(line, TokenSeparator)
				switch strings.TrimSpace(tokens[0]) {
				case "Starting items":
					startingItems := strings.Split(tokens[1], ", ")
					monkey.Items = lo.Map(startingItems, func(s string, _ int) int {
						item, _ := strconv.Atoi(strings.TrimSpace(s))
						return item
					})
				case "Operation":
					monkey.Operation = parse(tokens[1])
				case "Test":
					tk := strings.Split(strings.TrimSpace(tokens[1]), " ")[2]
					modulus, _ := strconv.Atoi(tk)
					monkey.TestModulus = modulus
				case "If true":
					target, _ := strconv.Atoi(strings.TrimSpace(strings.Split(tokens[1], "monkey")[1]))
					monkey.TargetTrue = target
				case "If false":
					target, _ := strconv.Atoi(strings.TrimSpace(strings.Split(tokens[1], "monkey")[1]))
					monkey.TargetFalse = target
				}
			}
			result = append(result, monkey)
		}
	}
	return
}

func parse(input string) Operation {
	poorMansLexerTokens := strings.Split(strings.TrimSpace(input), " ")
	var parserResult []string

	for _, token := range poorMansLexerTokens {
		switch strings.TrimSpace(token) {
		case "new":
			parserResult = append(parserResult, tokens.VARIABLE_NEW)
		case "old":
			parserResult = append(parserResult, tokens.VARIABLE_OLD)
		case "=":
			parserResult = append(parserResult, tokens.ASSIGN)
		case "*":
			parserResult = append(parserResult, tokens.MULTIPLY)
		case "+":
			parserResult = append(parserResult, tokens.ADD)
		default:
			parserResult = append(parserResult, strings.TrimSpace(token))
		}
	}

	operand, _ := strconv.Atoi(strings.TrimSpace(parserResult[4]))
	switch strings.TrimSpace(parserResult[3]) {
	case "+":
		return Add{operand}
	case "*":
		if parserResult[4] == tokens.VARIABLE_OLD {
			return Pow2{0}
		} else {
			return Mul{operand}
		}

	}

	return Add{0}
}