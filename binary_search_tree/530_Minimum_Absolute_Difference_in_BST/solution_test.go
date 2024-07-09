package main

import "testing"

// minimum absolute difference between the values
func Test_getMinimumDifference(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		tree := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 6},
		}
		output := getMinimumDifference(tree)
		if output != 1 {
			t.Errorf("Got %d; expected %d", output, 1)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val:   48,
				Left:  &TreeNode{Val: 12},
				Right: &TreeNode{Val: 49},
			},
		}
		output := getMinimumDifference(tree)
		if output != 1 {
			t.Errorf("Got %d; expected %d", output, 1)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 2},
			},
		}
		output := getMinimumDifference(tree)
		if output != 1 {
			t.Errorf("Got %d; expected %d", output, 1)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 3},
			},
		}
		output := getMinimumDifference(tree)
		if output != 2 {
			t.Errorf("Got %d; expected %d", output, 2)
		}
	})
}
