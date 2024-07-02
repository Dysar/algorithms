package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// O(n) is if we visit every node. All nodes are completely filled.
	// They have left node 100%

	// it will take log(n) time to calculate height. In worst case -> lg(n)*(lg(n)

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// calculate LH, RH
		lh := 0 // left height -> number of left nodes
		{
			leftNode := node.Left
			for leftNode != nil {
				lh++
				leftNode = leftNode.Left
			}
		}

		rh := 0 // right height -> number of right nodes
		{
			rightNode := node.Right
			if rightNode != nil {
				rh++
				rightNode = rightNode.Right
			}
		}

		if lh == rh {
			if lh == 0 {
				return 1
			}
			numOfNodes := math.Pow(2, float64(lh+1)) - 1
			return int(numOfNodes)
		} else {
			// if they are not equal ->
			//	go to left -> calculate LH, RH. return number
			leftNodeHeight := traverse(node.Left)
			//  go to right -> calculate LH, RH return number
			rightNodeHeight := traverse(node.Right)

			return leftNodeHeight + rightNodeHeight + 1
		}
	}

	if root == nil {
		return 0
	}
	return traverse(root)
}
