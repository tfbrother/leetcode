package _48__Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 8 ms, faster than 99.22% of Go online submissions for Sort List.
Memory Usage: 5 MB, less than 97.22% of Go online submissions for Sort List.
*/
// 排序链表，因为链表无法像数组一样索引，所以用快速排序很难。
// 用归并排序比较相对简单，只需要找到链表的中间结点即可。
// 第一次实现的时候，用例[4,2,1,3]输出为[1]，思考是哪里出错了。
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 计算中间结点
	head1, head2 := partionList(head)
	l, r := sortList(head1), sortList(head2)
	return mergeList(l, r)
}

// 将链表从中间一分为二
func partionList(head *ListNode) (head1, head2 *ListNode) {
	prev, slow, fast := head, head, head

	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}

	head1, head2, prev.Next = head, slow, nil
	return
}

// 合并两个有序链表，[21题]
func mergeList(head1, head2 *ListNode) *ListNode {
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
