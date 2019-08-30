package _9

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 71.43% of Go online submissions for Remove Nth Node From End of List.
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	for i := 1; i < n; i++ {
		// 给定的 n 保证是有效的。所以这个可以不用检测
		// 即链表的长度小于n，此时n无效，直接返回
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next
	}

	// slow指向的就是倒数第N个节点
	// TODO 第一版错误实现，就没考虑删除的是第一个节点这种情况
	if slow == head { // 特殊处理，此时prev==nil
		return head.Next
	}

	prev.Next = prev.Next.Next
	return head
}

// 主要是搞明白slow和fast的含义
// 此处slow直线的是倒数第n+1个节点
func removeNthFromEnd_cn(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
