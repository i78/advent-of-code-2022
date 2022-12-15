package signals

import (
	"github.com/samber/lo"
)

type Pair struct {
	Left  *Node
	Right *Node
}

type Node struct {
	parent   *Node
	value    *int
	children []*Node
}

func (n *Node) isLeaf() bool {
	return len(n.children) == 0
}

func (p *Pair) InOrder() bool {
	return Compare(*p) <= 0
}

func Compare(pair Pair) int {
	leftChildCount, rightChildCount := len(pair.Left.children), len(pair.Right.children)

	for pos := 0; pos < rightChildCount && pos < leftChildCount; pos++ {
		leftChild, rightChild := pair.Left.children[pos], pair.Right.children[pos]
		if leftChild.isLeaf() && rightChild.isLeaf() {
			if leftChild.value != nil && rightChild.value != nil && *(leftChild.value) != *(rightChild.value) {
				return *(leftChild.value) - *(rightChild.value)
			}
		} else {
			next := Compare(Pair{EmbedLeafInNode(leftChild), EmbedLeafInNode(rightChild)})
			if next != 0 {
				return next
			}
		}
	}
	return leftChildCount - rightChildCount
}

func EmbedLeafInNode(node *Node) *Node {
	return lo.Ternary(node.isLeaf() && node.value != nil,
		&Node{
			parent: node,
			children: []*Node{{
				parent:   node,
				value:    node.value,
				children: node.children,
			}},
		},
		node)
}
