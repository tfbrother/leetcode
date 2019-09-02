package _43

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 8 ms, faster than 94.04% of Go online submissions for Reorder List.
Memory Usage: 5.4 MB, less than 100.00% of Go online submissions for Reorder List.
*/
// 和[61][234]基本一摸一样
// 分割与合并链表两个子操作
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	head, head2 := partionList(head)
	head2 = reverseList(head2)
	node1, next1, next2 := head, head, head2

	for head2 != nil {
		next1, next2 = node1.Next, head2.Next

		node1.Next, head2.Next = head2, next1
		node1, head2 = next1, next2
	}

	return
}

func partionList(head *ListNode) (head1, head2 *ListNode) {
	prev, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}

	if fast != nil { // 奇数
		prev, slow = slow, slow.Next
	}

	head1, head2, prev.Next = head, slow, nil
	return
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur = head

	for cur != nil {
		next = cur.Next

		cur.Next, prev, cur = prev, cur, next
	}

	return prev
}
