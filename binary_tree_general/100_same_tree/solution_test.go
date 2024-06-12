package main

import "testing"

func Test_isSameTree(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {

		p := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		q := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		if !isSameTree(p, q) {
			t.Errorf("%v != %v", p, q)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		p := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		q := &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		}

		if same := isSameTree(p, q); same {
			t.Errorf("%v != %v", p, q)
		}
	})
}
