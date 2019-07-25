package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 8 ms, faster than 91.25% of Go online submissions for Add Two Numbers.
Memory Usage: 4.7 MB, less than 98.00% of Go online submissions for Add Two Numbers.
*/
// 解法一：先依次将两个链表的各结点进行相加，不处理进位。最后再循环新的链表处理进位。
// 没有说不允许修改链表，所以可以以l1作为头结点进行返回。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, l1}
	var cur *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			// 按位相加，进位后面的循环中来进行处理。
			l1.Val += l2.Val
			cur = l1
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil {
			cur.Next = l2
			break
		} else { // l2.Next == nil
			cur.Next = l1
			break
		}
	}

	// 处理进位的问题
	var prod int // 商
	cur = dummyHead.Next
	for cur != nil {
		//fmt.Println(cur.Val)
		if cur.Val >= 10 { // 处理进位
			prod = cur.Val / 10
			cur.Val = cur.Val % 10
			if cur.Next == nil {
				cur.Next = &ListNode{prod, nil}
			} else {
				cur.Next.Val += prod
			}
		}

		cur = cur.Next
	}

	return dummyHead.Next
}

/*
Runtime: 4 ms, faster than 99.18% of Go online submissions for Add Two Numbers.
Memory Usage: 4.7 MB, less than 98.00% of Go online submissions for Add Two Numbers.
*/
// 解法二：在解法一的基础上进行优化，相加时进行进位处理。较少一次循环。
// O(2N)优化成O(N)
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead, prod := &ListNode{0, l1}, 0
	var cur *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			l1.Val += l2.Val
			cur = l1
			l1, l2 = l1.Next, l2.Next
		} else if l1 == nil {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else { // l2.Next == nil
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
		//fmt.Println(cur.Val)
		if cur.Val >= 10 { // 处理进位
			prod = cur.Val / 10
			cur.Val = cur.Val % 10
			// TODO 这里的逻辑是核心
			if cur.Next == nil {
				// 这个新插入的结点要插入到l1中去才行，如果插入l2中将无法通过，为什么？
				// 因为当cur.Next == nil时，l1肯定为nil，而l2不一定为nil。仔细阅读自己的代码逻辑。所以此处不能放入到l2中。
				l1 = &ListNode{prod, nil}
				cur.Next = l1 // 如果不执行这段代码，将cur串起来，将不会得到正确的结果
			} else {
				cur.Next.Val += prod
			}
		}
	}

	return dummyHead.Next
}

// 解法三：不修改l1和l2
// 核心是carry的处理。
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
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
