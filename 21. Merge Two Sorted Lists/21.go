package _1__Merge_Two_Sorted_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.
*/
func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			prev.Next = head1
			head1 = head1.Next
		} else {
			prev.Next = head2
			head2 = head2.Next
		}

		prev = prev.Next
	}

	if head1 == nil {
		head1 = head2
	}

	prev.Next = head1

	return dummyHead.Next
}
