package fstree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllDirectories(t *testing.T) {
	fakeTree :=
		Node{Name: "/", Type: Directory, Size: 0, Children: []*Node{
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

	t.Run("Should find all directories", func(t *testing.T) {
		dirs := fakeTree.FindAllDirectories()
		assert.Len(t, dirs, 4)
	})
}

func TestGetDfhDirSize(t *testing.T) {
	testcases := []struct {
		dir          Node
		expectedSize int
	}{
		{
			dir: Node{Name: "e", Type: Directory, Size: 0, Children: []*Node{
				{Name: "i", Type: File, Size: 584, Children: nil},
			}},
			expectedSize: 584,
		}, {
			dir: Node{Name: "a", Type: Directory, Size: 0, Children: []*Node{
				{Name: "e", Type: Directory, Size: 0, Children: []*Node{
					{Name: "i", Type: File, Size: 584, Children: nil},
				}},
				{Name: "f", Type: File, Size: 29116, Children: nil},
				{Name: "g", Type: File, Size: 2557, Children: nil},
				{Name: "h.lst", Type: File, Size: 62596, Children: nil},
			}},
			expectedSize: 94853,
		}, {
			dir: Node{Name: "d", Type: Directory, Size: 0, Children: []*Node{
				{Name: "j", Type: File, Size: 4060174, Children: nil},
				{Name: "d.log", Type: File, Size: 8033020, Children: nil},
				{Name: "d.ext", Type: File, Size: 5626152, Children: nil},
				{Name: "k", Type: File, Size: 7214296, Children: nil},
			}},
			expectedSize: 24933642,
		},
	}

	for _, testcase := range testcases {
		t.Run("should return expected size", func(t *testing.T) {
			size := testcase.dir.DirectorySize()
			assert.Equal(t, testcase.expectedSize, size)
		})
	}
}
