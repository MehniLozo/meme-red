package set

import (
	hashmap "github.com/MehniLozo/meme-red/structs/hashmap"
	tree "github.com/MehniLozo/meme-red/structs/tree"
	types "github.com/MehniLozo/meme-red/structs/types"
)

type Set[T types.Flexy] struct {
	Tr   *tree.Tree[T]
	HMAP hashmap.HashMap
}
type NodeSet[T types.Flexy] struct {
	Tr    tree.Tree[T]
	HNode hashmap.HashMap
	Score float64
	Len   uint
	Name  []byte
}

func (s *Set[T]) newNode(name []byte, len uint, score float64) {
	node := &NodeSet[T]{
		Tr:    &tree.Tree[T]{},
		HNode: nil,
		Score: score,
		Len:   len,
	}
}
