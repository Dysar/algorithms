package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

	// TC: O(n), n = number of nodes, SC: O(h). h = height
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	// compare the values of the nodes
	return left.Val == right.Val &&
		// compare the outer subtrees if values are same
		isMirror(left.Left, right.Right) &&
		// compare the inner subtrees if outer subtrees are ok and values are same
		isMirror(left.Right, right.Left)
}
