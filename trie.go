package trie

import (
	"fmt"
	"log"
	"strings"
)

func NewTrie(value string) *Trie {
	return &Trie{
		Value: value,
	}
}

type Trie struct {
	Value    string
	Children []*Trie
}

// find a path in tree
func (me *Trie) Find(v string) *Trie {
	var cur *Trie
	parts := strings.Split(v, "/")
	index := 0
	for index = 0; index < len(parts); index++ {
		if cur == nil {
			break
		}
		c := cur.FindChild(parts[index])
		if c == nil {
			break
		}
		cur = c
		index++
	}
	return cur
}

func (me *Trie) FindChild(v string) *Trie {
	if me.Children == nil {
		return nil
	}
	for _, c := range me.Children {
		if c.Value == v {
			return c
		}
	}
	return nil
}

// insert an element into the tree
func (me *Trie) Insert(c *Trie) {
	me.Children = append(me.Children, c)
}

// insert a path 'p' into the tree, the boundries are path components
func (me *Trie) InsertPath(p string) *Trie {
	parts := strings.Split(p, "/")

	// find a root matching our prefix or allocate
	var (
		c   *Trie
		cur *Trie
	)
	cur = me
	for _, part := range parts {
		c = cur.FindChild(part)
		log.Printf("find part:%s return  %p", part, c)
		if c == nil {
			t := NewTrie(part)
			cur.Insert(t)
			cur = t
		} else {
			log.Printf("e found. Advancing")
			cur = c
		}
	}
	return cur
}

func (me *Trie) PrintTree() {
	fmt.Printf("%s\n", me.Value)
	for _, c := range me.Children {
		c.Print(1)
	}
}

func (me *Trie) Print(depth int) {
	pad := strings.Repeat(" ", depth)
	fmt.Printf("%s %s\n", pad, me.Value)
	for _, c := range me.Children {
		c.Print(depth + 1)
	}
}

type WalkFn func(depth int, path string, T *Trie)

func joinPath(a, b string) string {
	if a == "" {
		return b
	}
	return a + "/" + b
}

func (me *Trie) Walk(depth int, path string, fn WalkFn) {
	fn(depth, joinPath(path, me.Value), me)
	for _, c := range me.Children {
		c.Walk(depth+1, joinPath(me.Value, c.Value), fn)
	}
}
