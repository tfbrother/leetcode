package _06

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：迭代
func reverseList1(head *ListNode) *ListNode {
	var prev, cur, next *ListNode

	cur = head
	for cur != nil {
		next = cur.Next

		cur.Next, prev, cur = prev, cur, next
	}

	return prev
}

// 解法二：递归
func reverseList2_err(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := reverseList2_err(head.Next)
	prev.Next = head
	return prev
}

// 解法二：递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}
