package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := invertTree(root.Right) // O(n), size = O(n)
	r := invertTree(root.Left)  // O (n), size = O(n)
	root.Left = l
	root.Right = r
	return root

	// overall SC O(n), TC O(n)
}
