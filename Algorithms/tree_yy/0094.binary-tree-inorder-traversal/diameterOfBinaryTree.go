package Problem0094

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxDai int
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)
		maxDai = max(maxDai, leftDepth+rightDepth)
		return max(leftDepth, rightDepth) + 1
	}
	depth(root)
	return maxDai
}
