package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(v *TreeNode, maxLength *int) int {
	if v == nil {
		return -1
	}

	left := 1 + dfs(v.Left, maxLength)
	right := 1 + dfs(v.Right, maxLength)
	if *maxLength < left+right {
		*maxLength = left + right
	}

	if left < right {
		left = right
	}
	return left
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	dfs(root, &maxLength)
	return maxLength
}
