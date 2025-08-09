package Problem0094

import "fmt"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		fmt.Println("到达空节点，返回")
		return nil
	}
	fmt.Printf("进入节点 %d\n", root.Val)
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	fmt.Printf("交换节点 %d 的左右子树\n", root.Val)
	return root
}
