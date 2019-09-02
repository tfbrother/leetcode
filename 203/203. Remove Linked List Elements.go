package _03

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, cur := dummyHead, dummyHead.Next

	for cur != nil {
		if cur.Val == val {
			cur = cur.Next
			prev.Next = cur
		} else {
			prev.Next = cur
			cur, prev = cur.Next, prev.Next
		}

	}

	return dummyHead.Next
}
