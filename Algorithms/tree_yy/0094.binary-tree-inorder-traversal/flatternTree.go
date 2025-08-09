package Problem0094

import "fmt"

var prev *TreeNode // 全局指针，指向已构建链表的尾节点

func flatten(root *TreeNode) {
	if root == nil {
		fmt.Println("⚠️ 遇到空节点，返回")
		return
	}

	// 打印当前节点信息
	fmt.Printf("\n🔵 处理节点 %d\n", root.Val)
	if root.Left != nil || root.Right != nil {
		fmt.Printf("├─ 当前节点状态: 左=%v, 右=%v\n",
			getNodeVal(root.Left), getNodeVal(root.Right))
	}

	// 1. 先递归处理右子树（关键！）
	fmt.Printf("├─ 递归处理右子树【%v】\n", getNodeVal(root.Right))
	flatten(root.Right)

	// 2. 再递归处理左子树
	fmt.Printf("├─ 递归处理左子树【%v】\n", getNodeVal(root.Left))
	flatten(root.Left)

	// 3. 连接链表（核心操作）
	fmt.Printf("├─ 连接链表：")
	fmt.Printf("%d → %v\n", root.Val, getNodeVal(prev))
	root.Right = prev // 当前节点指向已构建的链表
	root.Left = nil   // 左指针置空

	// 4. 更新尾节点指针
	prev = root
	fmt.Printf("└─ 更新尾节点 → %d\n", root.Val)
}

// 辅助函数：安全获取节点值
func getNodeVal(node *TreeNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", node.Val)
}
