package main

import "testing"

func Test_maxDepth(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		node := &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}
		d := maxDepth(node)
		if d != 3 {
			t.Errorf("maxDepth() = %d, want %d", d, 3)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		node := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}
		d := maxDepth(node)
		if d != 2 {
			t.Errorf("maxDepth() = %d, want %d", d, 3)
		}
	})
}
