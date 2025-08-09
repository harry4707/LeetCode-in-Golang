package Problem0094

import "math"

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return valid(root, math.MinInt, math.MaxInt)

}

func valid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	} // 终止条件

	if node.Val < min || node.Val > max {
		return false
	}
	flag1 := valid(node.Left, min, node.Val)
	flag2 := valid(node.Right, node.Val, max)
	return flag1 && flag2

}
