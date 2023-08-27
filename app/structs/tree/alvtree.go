package tree

import (
	max "github.com/MehniLozo/meme-red/structs/ops"
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

func (t *Tree[T]) Put(key T, value T, p *Node[T], qp **Node[T]) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node[T]{key: key, val: value, parent: p}
		return true
	}

	// Unfinished business
	return false // for now
}

// TODO
/*func (t *Tree[T]) Push(k ...T) {
	for _, x := range k {
		t.Root = t.insert(t.Root, x)
	}
}*/
func (t *Tree[T]) height(root *Node[T]) int {
	if root == nil {
		return 1
	}
	var lHeight, rHeight int
	if root.children[0] != nil {
		lHeight = root.children[0].height
	}
	if root.children[1] != nil {
		rHeight = root.children[1].height
	}
	return 1 + max.Int(lHeight, rHeight)
}
func (t *Tree[T]) giveMeBalanceFact(root *Node[T]) int {
	var lHeight, rHeight int
	if root.children[0] != nil {
		lHeight = root.children[0].height
	}
	if root.children[1] != nil {
		rHeight = root.children[1].height
	}
	return lHeight - rHeight
}
