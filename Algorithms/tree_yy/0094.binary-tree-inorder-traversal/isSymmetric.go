package Problem0094

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Compare(root.Left, root.Right)
}

func Compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return Compare(left.Left, right.Left) && Compare(left.Right, right.Right)

}
