package main

import "testing"

func Test_invertTree(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		node := &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}

		inverted := invertTree(node)
		if inverted.Right.Val != 1 || inverted.Left.Val != 3 {
			t.Errorf("\ninverted right = %d, inverted left = %d", inverted.Right.Val, inverted.Left.Val)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		node := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 9},
			},
		}

		inverted := invertTree(node)
		if inverted.Right.Right.Val != 1 {
			t.Errorf("want %d, got %d", 1, inverted.Right.Right.Val)
		}
	})
}
