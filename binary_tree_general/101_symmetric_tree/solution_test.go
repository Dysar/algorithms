package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {

		node := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		}

		if !isSymmetric(node) {
			t.Errorf("isSymmetric(%v) = false, want true", node)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		node := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		}
		if isSymmetric(node) {
			t.Errorf("isSymmetric(%v) = true, want false", node)
		}
	})
}
