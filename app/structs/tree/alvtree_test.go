package tree

import (
	"testing"
)

func TestPush(t *testing.T) {
	t.Run("Left Left Rotation Testing", func(t *testing.T) {
		tree := &Tree[int]{
			Root: nil,
		}
		tree.Push(5, 4, 3)
		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key())
		}

		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height())
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3, not %d \n", root.Left().Key())
		}
		if root.Left().Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root != nil && root.Right() != nil && root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}

		if root != nil && root.Right() != nil && root.Right().height != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("Left Right Rotation Testing", func(t *testing.T) {
		tree := &Tree[int]{
			Root: nil,
		}
		//tree.Push(5, 3, 4)
		tree.Push(5, 6)
		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key())
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height())
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

}

/*

// Based on the instrusive Data structure pattern
type Container struct {
	root Node[uint32]
}

func addCont(c *Container, val uint32) {
	data := Node[uint32]{
			key: val,
			parent: nil,
			children: [2]*Node[uint32]{nil,nil},
			height : 1,
	}
	if c.root == nil {
		c.root = data
		return
	}

	cur := c.root

	for {
		var from Node[uint32]
		if

	}

}
*/
