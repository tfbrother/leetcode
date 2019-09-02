package _34

type ListNode struct {
	Val  int
	Next *ListNode
}

// 主要是满足O(1)空间复杂度
// 将链表一分为二，然后反转第二个链表，最后判断是否是回文链表
// 核心是节点个数奇偶情况的处理，
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	head1, head2 := partionList(head)
	head2 = reverseList(head2)

	for head1 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1, head2 = head1.Next, head2.Next
	}

	return true
}

func partionList(head *ListNode) (head1, head2 *ListNode) {
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}

	if fast != nil { // 说明是奇数个
		slow = slow.Next
	}
	head1, head2, prev.Next = head, slow, nil
	return
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur = head

	for cur != nil {
		next = cur.Next

		cur.Next = prev
		prev, cur = cur, next
	}

	return prev
}
