package test

import (
	"fmt"
	"leetcode-go/problem"
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	result := problem.MaxLevelSum(&problem.TreeNode{
		Val: 1,
		Left: &problem.TreeNode{
			Val: 2,
		},
		Right: &problem.TreeNode{
			Val: -3,
		},
	})
	fmt.Printf("result: %v\n", result)
}
