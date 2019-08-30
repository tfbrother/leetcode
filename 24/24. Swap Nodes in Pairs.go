package _4

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
Memory Usage: 2.1 MB, less than 66.67% of Go online submissions for Swap Nodes in Pairs.
*/
// 第一版的实现在链表中形成环了
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev, cur := dummyHead, dummyHead.Next

	var left, right, next *ListNode
	for cur != nil && cur.Next != nil {
		left, right, next = cur, cur.Next, cur.Next.Next
		prev.Next, right.Next = right, left
		left.Next = next // TODO 第一版的实现就没有加这句，导致在链表中形成环了
		prev, cur = left, next
	}

	return dummyHead.Next
}
