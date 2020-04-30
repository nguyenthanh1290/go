package bitree

import (
	"math"
)

var mps int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	mps = math.MinInt32

	pathSum(root)

	return mps
}

func pathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := max(0, pathSum(node.Left))
	right := max(0, pathSum(node.Right))

	mps = max(mps, left+right+node.Val)

	return max(left, right) + node.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
