package _6

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition List.
Memory Usage: 2.3 MB, less than 66.67% of Go online submissions for Partition List.
*/
// 解法一：
// 遍历一遍链表，将所有小于x节点放入链表1，其它放入链表2，最后将两个链表合并
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead1, dummyHead2 := &ListNode{}, &ListNode{}
	prev1, prev2, curNode := dummyHead1, dummyHead2, head

	for curNode != nil {
		next := curNode.Next
		curNode.Next = nil
		if curNode.Val >= x { // 放入链表2
			prev2.Next = curNode
			prev2 = prev2.Next
		} else { // 放入链表1
			prev1.Next = curNode
			prev1 = prev1.Next
		}

		curNode = next
	}

	prev1.Next = dummyHead2.Next
	return dummyHead1.Next
}

// 解法二：
// 遍历一遍链表，将所有大于x节点从原链表删除，然后放入新链表，最后将两个链表合并
// 删除头节点要特殊处理
func partition2(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyHead1, dummyHead2 := &ListNode{Next: head}, &ListNode{}
	prev1, prev2, curNode := dummyHead1, dummyHead2, head

	for curNode != nil {
		next := curNode.Next
		curNode.Next = nil

		if curNode.Val >= x { // 从链表1中删除，并放入链表2
			prev1.Next, prev2.Next = next, curNode
			prev2 = prev2.Next
		} else {
			prev1.Next = curNode
			prev1 = prev1.Next
		}

		curNode = next
	}

	prev1.Next = dummyHead2.Next
	return dummyHead1.Next
}
