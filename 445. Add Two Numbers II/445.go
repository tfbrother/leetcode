package _45__Add_Two_Numbers_II

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 8 ms, faster than 92.48% of Go online submissions for Add Two Numbers II.
Memory Usage: 4.9 MB, less than 80.95% of Go online submissions for Add Two Numbers II.
*/
// 解法一：
//  1.反转l1和l2。
//  2.参考[2]题实现，返回新链表head
//  3.反转head链表，并返回。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	head := add(l1, l2)
	return reverse(head)
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *ListNode
	cur = head

	// 注意递归结束条件
	for cur != nil {
		next = cur.Next // 暂存cur.Next
		cur.Next = prev // 进行反转
		prev = cur
		cur = next
	}

	return prev
}

// 两个链表相加，不能对l1、l2进行修改
func add(l1 *ListNode, l2 *ListNode) *ListNode {
	head, val, carry := &ListNode{}, 0, 0
	cur := head

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			val = l1.Val + l2.Val
			l1, l2 = l1.Next, l2.Next
		} else if l1 != nil {
			val = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			val = l2.Val
			l2 = l2.Next
		}

		// 处理上一次循环中产生的进位
		if carry == 1 {
			val += 1
			carry = 0
		}

		// 记录本一次循环中产生的进位
		if val >= 10 {
			val %= 10
			carry = 1
		}

		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	if carry == 1 {
		// 这里的Val只可能为1，要注意这个性质
		cur.Next = &ListNode{Val: 1}
	}

	return head.Next
}

/*
Runtime: 8 ms, faster than 92.48% of Go online submissions for Add Two Numbers II.
Memory Usage: 4.9 MB, less than 80.95% of Go online submissions for Add Two Numbers II.
*/
// 解法二：进阶
//  1.通过一次循环，找到两个链表的高度差。
//  2.依次将相应位的数进行相加，生成新的链表。不处理进位，否则有点麻烦。
//  3.循环新的链表，处理进位。
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2, diffLen := l1, l2, 0
	for cur1 != nil && cur2 != nil {
		cur1, cur2 = cur1.Next, cur2.Next
	}

	if cur1 == nil && cur2 != nil { // l1比l2短，交换
		l1, l2 = l2, l1
		cur1 = cur2
	}

	// l1不比l2短了。
	for cur1 != nil {
		diffLen++
		cur1 = cur1.Next
	}

	head := &ListNode{}
	cur, i := head, 0
	cur1, cur2 = l1, l2

	for i < diffLen {
		cur.Next = &ListNode{Val: cur1.Val}
		cur, cur1 = cur.Next, cur1.Next
		i++
	}

	for cur1 != nil {
		cur.Next = &ListNode{Val: cur1.Val + cur2.Val}
		cur, cur1, cur2 = cur.Next, cur1.Next, cur2.Next
	}

	// 核心：如何处理进位的问题
	head.Next = reverse(head.Next)
	prev, carry := head, 0
	for prev.Next != nil { // 核心，为何要比较prev.Next
		// 处理上一次循环中产生的进位信息
		if carry == 1 {
			prev.Next.Val++
			carry = 0
		}

		// 计算本次循环是否产生的进位
		if prev.Next.Val >= 10 {
			prev.Next.Val %= 10
			carry = 1
		}

		prev = prev.Next

	}

	if carry == 1 {
		// 核心，思考清楚逻辑
		prev.Next = &ListNode{Val: 1}
	}

	return reverse(head.Next)
}

// 解法三：利用栈
