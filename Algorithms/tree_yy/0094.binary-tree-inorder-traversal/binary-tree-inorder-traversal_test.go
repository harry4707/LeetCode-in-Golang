package Problem0094

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	pre []int
	in  []int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0094(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
			},
			ans{
				[]int{1, 3, 2},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, inorderTraversal(PreIn2Tree(p.pre, p.in)), "输入:%v", p)
	}
}

// PreIn2Tree 把 preorder 和 inorder 切片转换成 二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}
func Test_invertTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	root.Right = &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}
	invertTree(root)
}

func Test_Flattern(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println("=====================")
	fmt.Println("🌳 原始树结构 (先序遍历):")
	printTree(root)
	fmt.Println("\n=====================")

	prev = nil // 重置全局指针
	flatten(root)

	fmt.Println("\n=====================")
	printList(root)
	fmt.Println("=====================")
}

// 辅助函数：先序遍历打印树
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

// 打印展开后的链表
func printList(head *TreeNode) {
	fmt.Print("\n📋 最终链表结构: ")
	for head != nil {
		fmt.Printf("%d", head.Val)
		if head.Right != nil {
			fmt.Print(" → ")
		}
		head = head.Right
	}
	fmt.Println(" → nil")
}
