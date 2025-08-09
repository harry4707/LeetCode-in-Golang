package Problem0094

func KthSmallest(root *TreeNode, k int) int {
	count := 0
	res := -1

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || res != -1 {
			return
		}

		dfs(node.Left) // 遍历左子树
		count++        // 访问当前节点
		if count == k {
			res = node.Val
			return
		}

		dfs(node.Right)
	}
	dfs(root)
	return res

}
