package _5

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 4 ms, faster than 94.46% of Go online submissions for Reverse Nodes in k-Group.
Memory Usage: 3.6 MB, less than 100.00% of Go online submissions for Reverse Nodes in k-Group.
*/
// k=2时，就是第24题
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	cur := dummyHead.Next
	var tail *ListNode

	dummyHead.Next, tail = reverseKNode(cur, k)
	for tail != nil {
		cur = tail.Next
		if cur == nil {
			break
		}
		tail.Next, tail = reverseKNode(cur, k)
	}

	return dummyHead.Next
}

// 反转链表的前k个元素，并返回新的链表的首尾节点
func reverseKNode(head *ListNode, k int) (newHead, tail *ListNode) {
	var prev, cur, next *ListNode
	cur = head

	for i := 0; i < k; i++ {
		if cur == nil { // cur指向第k个节点，此时节点不足
			return head, nil
		}

		prev, cur = cur, cur.Next
	}

	cur, prev = head, nil
	for i := 0; i < k; i++ {
		next = cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	head.Next = cur
	return prev, head
}
