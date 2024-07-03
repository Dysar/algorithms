package main

import "testing"

func Test_averageOfLevels(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		tree := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}
		output := averageOfLevels(tree)
		if len(output) != 3 {
			t.Errorf("output was incorrect, got: %d, want: %d.", len(output), 3)
		}
		if output[0] != 3 || output[1] != 14.5 || output[2] != 11 {
			t.Errorf("output was incorrect, got: %f, want: %f", output, []float64{3, 14.5, 11})
		}
	})

	t.Run("case 2", func(t *testing.T) {
		tree := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val: 20,
			},
		}
		output := averageOfLevels(tree)
		if len(output) != 3 {
			t.Errorf("output was incorrect, got: %d, want: %d.", len(output), 3)
		}
		if output[0] != 3 || output[1] != 14.5 || output[2] != 11 {
			t.Errorf("output was incorrect, got: %f, want: %f", output, []float64{3, 14.5, 11})
		}
	})
}
