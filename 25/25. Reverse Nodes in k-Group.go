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

// 解法二：先求出链表的总长度，然后在进行反转
// 该解法更容易理解，实现。
func reverseKGroup2(head *ListNode, k int) *ListNode {
	var dummyNode, prev, cur, next *ListNode
	var listLen int
	cur = head

	// 循环遍历链表，获取链表的长度
	for cur != nil {
		cur = cur.Next
		listLen++
	}

	dummyNode = &ListNode{0, head}
	// 前一个node
	prev = dummyNode
	// 从cur开始翻转k个元素
	cur = dummyNode.Next

	// 翻转reverseNum次，每次翻转k个元素
	reverseNum := int(listLen / k)
	for i := 0; i < reverseNum; i++ {
		prev.Next, next = reverseList(cur, k)
		prev = cur
		cur = next
	}

	return dummyNode.Next
}

// 翻转链表的前k个元素，返回新链表的head以及第k+1个元素
func reverseList(head *ListNode, k int) (newHead, nextHead *ListNode) {
	var prev, cur, next *ListNode
	cur = head
	index := 0

	for index < k {
		next = cur.Next
		cur.Next = prev
		prev = cur

		cur = next
		index++
	}

	newHead = prev
	head.Next = cur
	nextHead = cur

	return
}
