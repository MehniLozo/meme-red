package tree

import (
	max "github.com/MehniLozo/meme-red/structs/ops"
	types "github.com/MehniLozo/meme-red/structs/types"
)

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

func (t *Tree[T]) Get(key T) (value interface{}, found bool) {
	node := t.GetNode(key)
	if node != nil {
		return node.val, true
	}
	return nil, false
}
func (t *Tree[T]) GetNode(key T) *Node[T] {
	r := t.Root
	for r != nil {
		a := t.Analog(key, r.key) // Look for a comparion analogy
		switch {
		case a == 0:
			return r
		case a < 0:
			r = r.children[0]
		case a > 0:
			r = r.children[1]
		}
	}
	return r
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

func (t *Tree[T]) lRotation(n *Node[T]) *Node[T] {
	a := n.children[1]
	al := a.children[0]
	a.children[0] = n
	n.children[1] = al

	if al != nil {
		al.parent = n
	}

	a.parent = n.parent
	n.parent = a

	n.height = t.height(n)
	a.height = t.height(a)

	return a
}
func (t *Tree[T]) rRotation(n *Node[T]) *Node[T] {
	a := n.children[0]
	ar := a.children[1]
	a.children[1] = n
	n.children[0] = ar

	if ar != nil {
		ar.parent = n
	}

	a.parent = n.parent
	n.parent = a

	n.height = t.height(n)
	a.height = t.height(a)

	return a
}

// When left subtree is too deep
func (Node[T]) fixLeft(t *Tree[T], r *Node[T]) *Node[T] {
	if r.children[0].Height() < r.children[1].Height() {
		r.children[0] = t.lRotation(r.children[0])
	}
	return t.rRotation(r)
}

func (Node[T]) fixRight(t *Tree[T], r *Node[T]) *Node[T] {
	if r.children[1].Height() < r.children[0].Height() {
		r.children[0] = t.rRotation(r.children[1])
	}
	return t.lRotation(r)
}

func min[T types.Flexy](node *Node[T]) *Node[T] {
	if node == nil {
		return node
	}

	for node.children[0] != nil {
		node = node.children[0]
	}

	return node
}

func maxi[T types.Flexy](node *Node[T]) *Node[T] {
	if node == nil {
		return node
	}

	for node.children[1] != nil {
		node = node.children[1]
	}

	return node
}
