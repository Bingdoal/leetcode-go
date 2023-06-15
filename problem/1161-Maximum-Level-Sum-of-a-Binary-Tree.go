package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxLevelSum(root *TreeNode) int {
	var (
		levelSum     = map[int]int{}
		nodeQueue    = []*TreeNode{}
		queueLevel   = []int{}
		queuePointer = 0
	)
	nodeQueue = append(nodeQueue, root)
	queueLevel = append(queueLevel, 0)
	for ; queuePointer < len(nodeQueue); queuePointer++ {
		node := nodeQueue[queuePointer]
		level := queueLevel[queuePointer]
		if node != nil {
			levelSum[level] += node.Val
			nodeQueue = append(nodeQueue,
				nodeQueue[queuePointer].Left,
				nodeQueue[queuePointer].Right)
			queueLevel = append(queueLevel, level+1, level+1)
		}
	}
	return getMaxLevel(levelSum)
}

func getMaxLevel(levelSum map[int]int) int {
	max := levelSum[0]
	maxLevel := 1
	for k, v := range levelSum {
		if v > max {
			max = v
			maxLevel = k + 1
		}
	}
	return maxLevel
}
