package Problem0094

func sortedArrayToBST(arr []int) *TreeNode {
	return buildBst(arr, 0, len(arr)-1)
}

func buildBst(nums []int, left, right int) *TreeNode {
	// 递归终止条件
	if left > right {
		return nil
	}

	//
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBst(nums, left, mid-1)
	root.Right = buildBst(nums, mid+1, right)

	return root

}
