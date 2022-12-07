package main

import (
	"dreese.de/aoc22/day7/fstree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAOCAssignments(t *testing.T) {
	fakeTree :=
		fstree.Node{Name: "/", Type: fstree.Directory, Size: 0, Children: []*fstree.Node{
			{Name: "a", Type: fstree.Directory, Size: 0, Children: []*fstree.Node{
				{Name: "e", Type: fstree.Directory, Size: 0, Children: []*fstree.Node{
					{Name: "i", Type: fstree.File, Size: 584, Children: nil},
				}},
				{Name: "f", Type: fstree.File, Size: 29116, Children: nil},
				{Name: "g", Type: fstree.File, Size: 2557, Children: nil},
				{Name: "h.lst", Type: fstree.File, Size: 62596, Children: nil},
			}},
			{Name: "b.txt", Type: fstree.File, Size: 14848514, Children: nil},
			{Name: "c.dat", Type: fstree.File, Size: 8504156, Children: nil},
			{Name: "d", Type: fstree.Directory, Size: 0, Children: []*fstree.Node{
				{Name: "j", Type: fstree.File, Size: 4060174, Children: nil},
				{Name: "d.log", Type: fstree.File, Size: 8033020, Children: nil},
				{Name: "d.ext", Type: fstree.File, Size: 5626152, Children: nil},
				{Name: "k", Type: fstree.File, Size: 7214296, Children: nil},
			}},
		}}
	allDirs := fakeTree.FindAllDirectories()

	t.Run("Part One", func(t *testing.T) {
		totalSmallDirectorySizes := totalSizeOfSmallDirectories(allDirs, 100000)
		assert.Equal(t, 95437, totalSmallDirectorySizes)
	})

	t.Run("Part Two", func(t *testing.T) {
		const size, required = 70000000, 30000000

		spaceToReclaim := spaceToBeReclaimed(size, required, fakeTree.DirectorySize())
		smallestViableDirectorySize := smallestDirectorySuitableForDeletion(allDirs, spaceToReclaim)

		assert.Equal(t, 24933642, smallestViableDirectorySize)
	})
}
