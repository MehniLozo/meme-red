package tree

import (
	types "github.com/MehniLozo/meme-red/structs/types"
)

type Analog func(x, y interface{}) int

type Node[T types.Flexy] struct {
	key      T
	val      interface{}
	parent   *Node[T]
	children [2]*Node[T]
	height   int
}

type Tree[T types.Flexy] struct {
	Root   *Node[T]
	Analog Analog
	size   int
}

func (t *Tree[T]) Empty() bool {
	return t.size == 0
}

func (t *Tree[T]) Size() int {
	return t.size
}

func (t *Tree[T]) GetNode(key T) {

}

func (n *Node[T]) Parent() Node[T] {
	return *n.parent
}

func (n *Node[T]) Key() T {
	return n.key
}

func (n *Node[T]) Left() Node[T] {
	return *n.children[0]
}

func (n *Node[T]) Right() Node[T] {
	return *n.children[1]
}

func (n *Node[T]) Height() int {
	return n.height
}
