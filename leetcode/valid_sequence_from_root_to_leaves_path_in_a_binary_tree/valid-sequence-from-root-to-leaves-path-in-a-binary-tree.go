package bitree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil {
		return len(arr) == 0
	}

	return isValid(root, arr, 0)
}

func isValid(node *TreeNode, arr []int, i int) bool {
	// invalid cases:
	// 1. val is not equal
	if node.Val != arr[i] {
		return false
	}

	// 2. not leaf node
	if i == len(arr)-1 {
		return node.Left == nil && node.Right == nil
	}

	// continue matching:
	if node.Left != nil && isValid(node.Left, arr, i+1) {
		return true
	}
	if node.Right != nil && isValid(node.Right, arr, i+1) {
		return true
	}

	return false
}
