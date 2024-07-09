package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// in BST left child is less, right is greater.
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// it's not solvable with traversal as we need to find the difference
	// between any 2 nodes in the tree

	var elements []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		elements = append(elements, node.Val)

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	sort.Ints(elements) // O(nlogn)

	var minDiff int
	for i := range elements {
		// if i = 1, len must be 3 as I need 2nd and 3rd elements
		if len(elements) < i+2 {
			return minDiff
		}
		diff := elements[i+1] - elements[i]
		if minDiff == 0 {
			minDiff = diff
		} else {
			minDiff = min(minDiff, diff)
		}
	}

	return minDiff
}
