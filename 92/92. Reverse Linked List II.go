package _2

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
Memory Usage: 2 MB, less than 75.00% of Go online submissions for Reverse Linked List II.
*/
// 将链表分成三段来理解
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	cur, prev1, i := dummyHead, dummyHead, 0

	for i = 0; i < m; i++ {
		prev1, cur = cur, cur.Next
	}

	end2 := cur // 暂存起来,就成为反转后的尾巴节点
	// prev1指向第m-1个节点，cur指向的就是第m个节点
	// 反转[m,n]个节点
	var prev2, next *ListNode
	for cur != nil && i <= n {
		i++
		next = cur.Next
		cur.Next = prev2

		prev2, cur = cur, next
	}

	// 反转后，prev2指向头节点，end2指向尾部节点，cur指向第三段的节点
	prev1.Next = prev2
	end2.Next = cur

	return dummyHead.Next
}
