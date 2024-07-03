package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// return the average value of the nodes on each level
func averageOfLevels(root *TreeNode) []float64 {
	var (
		averagesPerLevel = make([]float64, 0)
		nextLvlNodes     = make([]*TreeNode, 0)
		bfs              func(nodes ...*TreeNode)
	)
	bfs = func(nodes ...*TreeNode) {

		sum := 0
		nextLvlNodes = []*TreeNode{}

		// we traverse all nodes on a given level
		for _, node := range nodes {
			if node == nil {
				return
			}
			// accumulate the current level node values to calculate the average
			sum += node.Val

			if node.Left != nil {
				// if there is a child, add it to the next level nodes
				nextLvlNodes = append(nextLvlNodes, node.Left)
			}
			if node.Right != nil {
				// if there is a child, add it to the next level nodes
				nextLvlNodes = append(nextLvlNodes, node.Right)
			}
		}
		averagesPerLevel = append(averagesPerLevel, float64(sum)/float64(len(nodes)))
		if len(nextLvlNodes) != 0 {
			bfs(nextLvlNodes...)
		}
	}
	bfs(root)
	return averagesPerLevel
}
