package Problem0094

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)     // current queue length
		levelVals := make([]int, 0) //存储本层的节点值

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		result = append(result, levelVals)
	}
	return result
}
