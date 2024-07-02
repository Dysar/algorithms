package main

import "testing"

func Test_hasPathSum(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		tree := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		output := hasPathSum(tree, 5)
		if output != false {
			t.Errorf("expected: %v, got: %v", false, output)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		output := hasPathSum(nil, 0)
		if output != false {
			t.Errorf("expected: %v, got: %v", false, output)
		}
	})

}
