package fstree

import (
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const ItemSeparator = " "
const Root = "/"

type ParserState struct {
	CurrentPath []string
}

type ParserLineEvaluator func(line string, state *ParserState, tree *Node)

func NewFsTree(raw string) *Node {
	node := NewDirectoryNode(Root)
	parserState := ParserState{CurrentPath: []string{}}

	for _, logLine := range strings.Split(raw, LineSeparator) {
		ResolveLineEvaluationFn(logLine)(logLine, &parserState, node)
	}
	return node
}

func ResolveLineEvaluationFn(line string) ParserLineEvaluator {
	switch {
	case strings.HasPrefix(line, "$ cd"):
		return Chdir
	case strings.HasPrefix(line, "$ ls"):
		return NothingToDo
	default:
		return CommandOutput
	}
}

func CommandOutput(line string, state *ParserState, tree *Node) {
	if strings.HasPrefix(line, "dir ") {
		node := NewDirectoryNode(strings.Split(line, ItemSeparator)[1])
		tree.InsertAtPath(state.CurrentPath, node)
	} else {
		tokens := strings.Split(line, ItemSeparator)
		size, _ := strconv.Atoi(tokens[0])
		node := NewFileNode(tokens[1], size)
		tree.InsertAtPath(state.CurrentPath, node)
	}
}

func Chdir(line string, state *ParserState, _ *Node) {
	dirToken := strings.Split(line, ItemSeparator)[2]

	switch dirToken {
	case "/":
		state.CurrentPath = []string{}
	case "..":
		state.CurrentPath = state.CurrentPath[0 : len(state.CurrentPath)-1]
	default:
		state.CurrentPath = append(state.CurrentPath, dirToken)
	}
}

func NothingToDo(string, *ParserState, *Node) {}
