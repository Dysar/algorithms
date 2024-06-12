package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	}
	if q == nil {
		// p can't be nil here.
		return false
	}

	//both are not nil
	if p.Val != q.Val {
		return false
	}

	if leftSame := isSameTree(p.Left, q.Left); !leftSame {
		return false
	}
	if rightSame := isSameTree(p.Right, q.Right); !rightSame {
		return false
	}
	return true
}
