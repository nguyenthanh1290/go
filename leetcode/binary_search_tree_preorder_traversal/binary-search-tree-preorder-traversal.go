package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	right := len(preorder)

	for i := range preorder {
		if preorder[i] > val {
			right = i
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  bstFromPreorder(preorder[1:right]),
		Right: bstFromPreorder(preorder[right:]),
	}
}

// tags: tree
