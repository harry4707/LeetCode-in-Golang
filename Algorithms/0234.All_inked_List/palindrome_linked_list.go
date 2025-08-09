package _234_Palindrome_inked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true // 空链表或单节点必为回文
	}

	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reSlow := reverseList(slow)
	newHead := head

	for reSlow != nil {
		if reSlow.Val != newHead.Val {
			return false
		}
		reSlow = reSlow.Next
		newHead = newHead.Next
	}
	return true

}

func reverseList(head *ListNode) *ListNode {
	var preHead *ListNode
	curHead := head
	for curHead != nil {
		next := curHead.Next
		curHead.Next = preHead

		preHead = curHead
		curHead = next
	}
	return preHead
}
