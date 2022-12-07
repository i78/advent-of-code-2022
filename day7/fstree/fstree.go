package fstree

import (
	"github.com/samber/lo"
)

type NodeType byte

const (
	Directory NodeType = iota
	File
)

type Node struct {
	Name     string
	Type     NodeType
	Size     int
	Children []*Node
}

func NewDirectoryNode(path string) *Node {
	return &Node{
		Name:     path,
		Type:     Directory,
		Size:     0,
		Children: []*Node{},
	}
}

func NewFileNode(name string, size int) *Node {
	return &Node{
		Name:     name,
		Type:     File,
		Size:     size,
		Children: nil,
	}
}

func (node *Node) InsertAtPath(path []string, newNode *Node) {
	if len(path) == 0 {
		node.AddChild(newNode)
	} else {
		topmostItemName := path[0]
		targetNode, exists := node.FindChild(topmostItemName)

		if !exists {
			targetNode = node.AddChild(NewDirectoryNode(topmostItemName))
		}

		targetNode.InsertAtPath(path[1:], newNode)
	}
	return
}

func (node *Node) FindChild(path string) (result *Node, found bool) {
	return lo.Find(node.Children, func(it *Node) bool { return it.Name == path })
}

func (node *Node) AddChild(child *Node) *Node {
	if _, exists := node.FindChild(child.Name); !exists {
		node.Children = append(node.Children, child)
	}
	return child
}

func (node *Node) FindAllDirectories() []*Node {
	results := []*Node{node}
	for _, child := range node.Children {
		if child.Type == Directory {
			results = append(results, child.FindAllDirectories()...)
		}
	}
	return results
}

func (node *Node) DirectorySize() (result int) {
	for _, child := range node.Children {
		switch child.Type {
		case File:
			result += child.Size
		case Directory:
			result += child.DirectorySize()
		}
	}
	return
}
