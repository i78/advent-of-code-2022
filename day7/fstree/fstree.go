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
	Parent   *Node
}

func NewDirectoryNode(path string, parent *Node) *Node {
	return &Node{
		Name:     path,
		Type:     Directory,
		Size:     0,
		Children: make([]*Node, 0, 2),
		Parent:   parent,
	}
}

func NewFileNode(name string, size int, parent *Node) *Node {
	return &Node{
		Name:     name,
		Type:     File,
		Size:     size,
		Children: nil,
		Parent:   parent,
	}
}

func (node *Node) FindChild(path string) (result *Node, found bool) {
	return lo.Find(node.Children, func(it *Node) bool { return it.Name == path })
}

func (node *Node) AddChild(child *Node) *Node {
	node.Children = append(node.Children, child)
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
