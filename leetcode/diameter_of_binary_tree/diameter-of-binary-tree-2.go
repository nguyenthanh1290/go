func diameterOfBinaryTree(root *TreeNode) int {
	var diameter = 0
	calculateMaxHeight(root, &diameter)
	return diameter
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateMaxHeight(node *TreeNode, diameter *int) int {
	var maxHeight = 0

	if node != nil {
		var maxHeightLeftNode = calculateMaxHeight(node.Left, diameter)
		var maxHeightRightNode = calculateMaxHeight(node.Right, diameter)

		*diameter = max(*diameter, maxHeightLeftNode+maxHeightRightNode)
		maxHeight = max(maxHeightLeftNode, maxHeightRightNode) + 1
	}

	return maxHeight
}