package _3__Remove_Duplicates_from_Sorted_List

// 和26题一摸一样，只是对象由有序数组变成有序链表了。
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 4 ms, faster than 90.43% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil { // 为nil的情况单独处理
		return head
	}
	prev, cur := head, head.Next

	for cur != nil {
		if cur.Val == prev.Val { // 删除cur结点
			prev.Next = cur.Next
		} else { // 依次往下寻找
			prev = cur
		}

		cur = cur.Next
	}

	return head
}

// head为nil的情况就不用单独处理了。
func deleteDuplicates2(head *ListNode) *ListNode {
	prev := head

	for prev != nil && prev.Next != nil {
		if prev.Next.Val == prev.Val { // 删除prev.Next结点
			prev.Next = prev.Next.Next
		} else { // 依次往下寻找
			prev = prev.Next
		}
	}

	return head
}
