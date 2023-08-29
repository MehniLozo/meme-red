package tree

import (
	types "github.com/MehniLozo/meme-red/structs/types"
)

type position byte

type Iterator[T types.Flexy] struct {
	tree *Tree[T]
	node *Node[T]
	pos  position
}

const (
	dep, middle, final position = 0, 1, 2
)

//func (tree *Tree[T]) Iterator()

/* func (iterator *Iterator[T]) Next() bool {
	switch iterator.pos {
	case dep:
		iterator.pos = middle
		iterator.node = iterator.tree.
	}
}*/
