package fstree

import (
	"strconv"
	"strings"
)

const LineSeparator = "\n"
const ItemSeparator = " "
const Root = "/"

type ParserState struct {
	CurrentDirNode *Node
}

type ParserLineEvaluator func(line string, state *ParserState, root *Node)

func NewFsTree(raw string) *Node {
	node := NewDirectoryNode(Root, nil)
	parserState := ParserState{CurrentDirNode: node}

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

func CommandOutput(line string, state *ParserState, _ *Node) {
	if strings.HasPrefix(line, "dir ") {
		node := NewDirectoryNode(strings.Split(line, ItemSeparator)[1], state.CurrentDirNode)
		state.CurrentDirNode.AddChild(node)
	} else {
		tokens := strings.Split(line, ItemSeparator)
		size, _ := strconv.Atoi(tokens[0])
		node := NewFileNode(tokens[1], size, state.CurrentDirNode)
		state.CurrentDirNode.AddChild(node)
	}
}

func Chdir(line string, state *ParserState, root *Node) {
	dirToken := strings.SplitAfterN(line, ItemSeparator, 3)[2]

	switch dirToken {
	case "/":
		state.CurrentDirNode = root
	case "..":
		state.CurrentDirNode = state.CurrentDirNode.Parent
	default:
		nextNode, _ := state.CurrentDirNode.FindChild(dirToken)
		state.CurrentDirNode = nextNode
	}
}

func NothingToDo(string, *ParserState, *Node) {}
