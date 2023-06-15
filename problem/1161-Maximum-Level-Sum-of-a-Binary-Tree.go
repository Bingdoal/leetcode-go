package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxLevelSum(root *TreeNode) int {
	var (
		nodeQueue = []*TreeNode{}
		level     = 1
		max       = root.Val
		maxLevel  = 1
	)
	nodeQueue = append(nodeQueue, root)

	for len(nodeQueue) > 0 {
		currentLevel := make([]*TreeNode, len(nodeQueue))
		copy(currentLevel, nodeQueue)
		nodeQueue = []*TreeNode{}

		levelSum := 0
		notNullNode := 0
		for _, node := range currentLevel {
			if node != nil {
				notNullNode++
				levelSum += node.Val
				nodeQueue = append(nodeQueue, node.Left, node.Right)
			}
		}
		if notNullNode > 0 && levelSum > max {
			max = levelSum
			maxLevel = level
		}
		level++
	}
	return maxLevel
}
