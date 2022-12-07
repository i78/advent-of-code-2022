package fstree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseShellHistory(t *testing.T) {

	t.Run("should return single directory as expected", func(t *testing.T) {
		const singleDirectory = "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d"

		result := NewFsTree(singleDirectory)

		expectedResult := &Node{Name: "/", Type: Directory, Size: 0, Children: []*Node{
			{Name: "a", Type: Directory, Size: 0, Children: []*Node{}},
			{Name: "b.txt", Type: File, Size: 14848514, Children: nil},
			{Name: "c.dat", Type: File, Size: 8504156, Children: nil},
			{Name: "d", Type: Directory, Size: 0, Children: []*Node{}},
		}}

		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return nested directory as expected", func(t *testing.T) {
		const singleDirectory = "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k"

		result := NewFsTree(singleDirectory)

		expectedResult :=
			&Node{Name: "/", Type: Directory, Size: 0, Children: []*Node{
				{Name: "a", Type: Directory, Size: 0, Children: []*Node{
					{Name: "e", Type: Directory, Size: 0, Children: []*Node{
						{Name: "i", Type: File, Size: 584, Children: nil},
					}},
					{Name: "f", Type: File, Size: 29116, Children: nil},
					{Name: "g", Type: File, Size: 2557, Children: nil},
					{Name: "h.lst", Type: File, Size: 62596, Children: nil},
				}},
				{Name: "b.txt", Type: File, Size: 14848514, Children: nil},
				{Name: "c.dat", Type: File, Size: 8504156, Children: nil},
				{Name: "d", Type: Directory, Size: 0, Children: []*Node{
					{Name: "j", Type: File, Size: 4060174, Children: nil},
					{Name: "d.log", Type: File, Size: 8033020, Children: nil},
					{Name: "d.ext", Type: File, Size: 5626152, Children: nil},
					{Name: "k", Type: File, Size: 7214296, Children: nil},
				}},
			}}

		assert.Equal(t, expectedResult, result)
	})
}
