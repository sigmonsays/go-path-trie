package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {

	pt := NewTrie("(root)")

	// insert parths
	paths := []string{
		"a/b",
		"a/b/foo",
		"a/b/bar",
		"d/e",
	}

	for _, path := range paths {
		pt.InsertPath(path)
	}
	pt.PrintTree()

	// find common prefix by scanning nodes with >1
	for _, c := range pt.Children {

		if len(c.Children) > 1 {
			c.Walk(0, "", func(depth int, path string, c *Trie) {
				if len(c.Children) > 1 {
					fmt.Printf("Common prefix: %s\n", path)
				}
			})
		}
	}

}
