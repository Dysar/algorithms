package main

import "testing"

func Test_countNodes(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := countNodes(nil)
		if output != 0 {
			t.Errorf("expected: %d, got: %d", 0, output)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := countNodes(&TreeNode{Val: 1})
		if output != 1 {
			t.Errorf("expected: %d, got: %d", 1, output)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 6},
			},
		}
		output := countNodes(tree)
		if output != 6 {
			t.Errorf("expected: %d, got: %d", 6, output)
		}
	})
}
