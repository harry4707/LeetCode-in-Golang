package Problem0094

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				res = append(res, node.Val) //处理右子树节点
			}
			if node.Left != nil {
				queue = append(queue, node.Left) //left
			}

			if node.Right != nil {
				queue = append(queue, node.Right) // left
			}

		}
	}
	return res
}
