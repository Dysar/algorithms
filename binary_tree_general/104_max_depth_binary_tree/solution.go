package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 */
func maxDepth(root *TreeNode) int {
	// brute force: traverse all.
	// Look at the node, if left found, look at that one. if right found, look at that.
	//
	if root == nil {
		return 0
	}

	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	return max(maxLeft, maxRight) + 1
}
