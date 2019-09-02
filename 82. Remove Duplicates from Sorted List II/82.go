package _2__Remove_Duplicates_from_Sorted_List_II

// 和26题较相似，只是对象由有序数组变成有序链表了。
// 和83题一样
type ListNode struct {
	Val  int
	Next *ListNode
}

// 该版本有两个错误：
// 1. 不应该把 cur.Next != nil 放入for 循环的条件中，导致的结果就是[1,1,1]用例失败了，输出的是[1]，期望的是[]。
// 2. for循环中的第一个else，不应该用prev != dummyHead，而应该用 cur != head，因为增加该条件的目的是避免
// 	  [0]这种情况被错误处理，如果golang中也有类似整型的nil的话，prevVal的初始值设置成nil，就需要这个条件了。
func deleteDuplicates1_err(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{0, head}
	// 这里是核心：因为要删除结点，所以必须要知道该结点的prev结点才能删除该结点。
	// 所以prev初始值设置为dummyHead，因为head也可能被删除的。
	prev, cur, prevVal := dummyHead, head, head.Val

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val { // 删除两个结点
			prevVal = cur.Val   // 这行代码没有写，导致用例[1,2,2,2]失败。
			cur = cur.Next.Next // 删除两个结点，cur和cur.Next
			prev.Next = cur
		} else if cur.Val == prevVal && prev != dummyHead { // TODO 删除当前结点，核心
			// TODO [1,1,1]用例失败了，输出的是[1]，期望的是[]。
			cur = cur.Next // 删除cur结点
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
			prevVal = prev.Val
		}
	}

	return dummyHead.Next
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List II.
Memory Usage: 2.9 MB, less than 97.56% of Go online submissions for Remove Duplicates from Sorted List II.
*/
// 完全借鉴26题的解法。
// 要删除所有重复值的结点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{0, head}
	// 这里是核心：因为要删除结点，所以必须要知道该结点的prev结点才能删除该结点。
	// 所以prev初始值设置为dummyHead，因为head也可能被删除的。
	// TODO prevVal的含义：记录上一个值。类似26题中的k的含义。此处设计是该题的核心，要理解透彻。
	prev, cur, prevVal := dummyHead, head, head.Val

	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val { // 删除两个结点
			prevVal = cur.Val   // 这行代码没有写，导致用例[1,2,2,2]失败。
			cur = cur.Next.Next // 删除两个结点，cur和cur.Next
			prev.Next = cur
		} else if cur.Val == prevVal && cur != head { // 删除当前结点
			// TODO [1,1,1]用例失败了，输出的是[1]，期望的是[]。
			// [1,2,2,2]，输出的是[1,2]，期望的是[1]。
			cur = cur.Next // 删除cur结点
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
			prevVal = prev.Val
		}
	}

	return dummyHead.Next
}

// 这个写法借鉴别人的提交，代码实现优美，思想简单，值得借鉴。
func deleteDuplicates2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}

	// 从链表头部之前开始
	prev := dummyHead
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val { // 该循环执行完，head指向重复元素的最后一个。
			head = head.Next
		}

		if prev.Next == head {
			prev = prev.Next
		} else { // 表示head跳过了重复节点，此时不移动prev
			prev.Next = head.Next
		}

		head = head.Next
	}

	return dummyHead.Next
}
