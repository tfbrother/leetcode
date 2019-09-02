package _47

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(N^2)
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead1, dummyHead2 := &ListNode{Next: head}, &ListNode{Next: head.Next}
	var prev1, prev2, node1, node2, next *ListNode
	prev2, node2, head.Next = dummyHead2, dummyHead2.Next, nil

	for node2 != nil {
		prev1, node1 = dummyHead1, dummyHead1.Next
		for node1 != nil {
			if node1.Val >= node2.Val { // node2插入到node1的前面
				next = node2.Next
				prev2.Next = node2.Next // 链表2中删除node2

				prev1.Next = node2 // node2插入链表1
				node2.Next = node1

				node2 = next
				break
			} else { // 继续往后找
				prev1, node1 = node1, node1.Next
			}
		}

		if node1 == nil { // node2插入链表1的末尾
			next = node2.Next
			prev2.Next = node2.Next // 链表2中删除node2

			prev1.Next = node2 // node2插入链表1
			node2.Next = node1

			node2 = next
		}
	}

	return dummyHead1.Next
}
