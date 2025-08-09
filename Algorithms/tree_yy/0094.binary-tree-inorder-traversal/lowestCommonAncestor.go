package Problem0094

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root // p 和 q 分居两侧，当前节点为 LCA
	}
	if left != nil {
		return left // LCA 在左子树
	}
	return right // LCA 在右子树

}
