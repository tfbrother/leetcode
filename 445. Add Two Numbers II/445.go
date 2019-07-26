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

// 解法二的基础上进行了优化，第三步不用反转链表就能处理进位。
func addTwoNumbers21(l1 *ListNode, l2 *ListNode) *ListNode {
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
	cur, carry := head, 1
	for carry == 1 {
		cur, carry = head.Next, 0
		for cur != nil && cur.Next != nil { // 核心，为何要比较prev.Next
			if cur.Next.Val > 9 {
				cur.Val++
				carry = 1
				cur.Next.Val %= 10
			}
			cur = cur.Next
		}
	}

	// 需要在链表头增加一个链表
	if head.Next.Val > 9 {
		// 核心，思考清楚逻辑
		prev, next := &ListNode{Val: 1}, head.Next
		head.Next = prev
		next.Val %= 10
		prev.Next = next
	}

	return head.Next
}

/*
Runtime: 4 ms, faster than 98.21% of Go online submissions for Add Two Numbers II.
Memory Usage: 5.7 MB, less than 31.58% of Go online submissions for Add Two Numbers II.
*/
// 解法三：利用两个栈
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2, cur1, cur2 := make([]int, 0), make([]int, 0), l1, l2

	for cur1 != nil {
		stack1 = append(stack1, cur1.Val)
		cur1 = cur1.Next
	}

	for cur2 != nil {
		stack2 = append(stack2, cur2.Val)
		cur2 = cur2.Next
	}

	if len(stack1) < len(stack2) {
		stack1, stack2 = stack2, stack1
	}

	i, j, k, carry, val, dummyHead := 1, 0, 0, 0, 0, &ListNode{}
	cur := dummyHead

	// 对两个栈进行相加
	for ; i <= len(stack1); i++ {
		j, k = len(stack1)-i, len(stack2)-i
		if k < 0 {
			val = stack1[j]
		} else {
			val = stack1[j] + stack2[k]
		}

		if carry == 1 {
			val++
			carry = 0
		}

		if val >= 10 {
			val %= 10
			carry = 1
		}

		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	if carry == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return reverse(dummyHead.Next)
}

// 解法3.1：在解法三的基础上进行优化，不用反转链表
func addTwoNumbers31(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2, cur1, cur2 := make([]int, 0), make([]int, 0), l1, l2

	for cur1 != nil {
		stack1 = append(stack1, cur1.Val)
		cur1 = cur1.Next
	}

	for cur2 != nil {
		stack2 = append(stack2, cur2.Val)
		cur2 = cur2.Next
	}

	if len(stack1) < len(stack2) {
		stack1, stack2 = stack2, stack1
	}

	i, j, k, carry, val := 1, 0, 0, 0, 0
	var head, node *ListNode

	// 对两个栈进行相加
	for ; i <= len(stack1) || carry == 1; i++ {
		j, k, val = len(stack1)-i, len(stack2)-i, 0
		if j >= 0 {
			val += stack1[j]
		}
		if k >= 0 {
			val += stack2[k]
		}

		if carry == 1 {
			val++
			carry = 0
		}

		if val >= 10 {
			val %= 10
			carry = 1
		}

		// TODO 这里就是反转链表，或者说是从尾部创建链表。
		node = &ListNode{Val: val}
		node.Next = head
		head = node
	}

	return head
}

// 参考别人的提交
// 	用例[7,2,4,3]，[5,6,4]
// 	1.自己的解法二相似，只有解法二的第三步不一样，第三步实际上是处理链表[7,7,10,7]中的进位信息
//	2.该题第三步的解法是通过循环链表处理进位，而不需要反转链表再处理进位。仔细阅读该部分代码。
func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	len1 := 0
	c1 := l1
	len2 := 0
	c2 := l2

	for c1 != nil {
		len1++
		c1 = c1.Next
	}

	for c2 != nil {
		len2++
		c2 = c2.Next
	}

	if len1 < len2 {
		len1, len2 = len2, len1
		l1, l2 = l2, l1
	}

	c1 = l1
	for i := 0; i < len1-len2; i++ {
		c1 = c1.Next
	}

	c2 = l2
	for c1 != nil {
		c1.Val += c2.Val
		c1 = c1.Next
		c2 = c2.Next
	}

	// 处理进位，flag默认为有进位需要处理。
	flag := true
	for flag {
		flag = false
		c1 = l1 // 从头开始循环处理。
		for c1 != nil && c1.Next != nil {
			c1.Val += c1.Next.Val / 10
			if c1.Next.Val > 9 {
				flag = true // 产生了进位，需要再次从头开始循环链表处理。因为链表是倒序的。
			}
			c1.Next.Val %= 10
			c1 = c1.Next
		}
	}

	// 判断整体是否需要进位
	if l1.Val >= 10 {
		ret := ListNode{l1.Val / 10, l1}
		l1.Val %= 10
		ret.Next = l1
		return &ret
	} else {
		return l1
	}
}

// 参考别人的提交，和自己的解法三相似，只是代码似乎更精简一点。
// 	1.该解法更精妙的地方在于对两个栈进行相加时并没有像解法三那样最后需要反转链表。
//	2.搞清楚是如何在循环中进行反转的。
func addTwoNumbers5(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)

	stacking(l1, &stack1)
	stacking(l2, &stack2)
	sp1, sp2 := len(stack1), len(stack2)

	var head *ListNode
	val, carry := 0, 0

	// 对两个栈进行相加
	for sp1 != 0 || sp2 != 0 || carry != 0 {
		val = 0
		if sp1 != 0 {
			sp1--
			val += stack1[sp1]
		}
		if sp2 != 0 {
			sp2--
			val += stack2[sp2]
		}

		val += carry
		carry = val / 10
		// TODO 仔细和反转链表代码对比，这他妈其实就是在反转链表啊。
		// 本质就是从头开始创建链表或者从尾部开始创建链表的区别。这是基本功啊。
		node := &ListNode{val % 10, nil}
		node.Next = head
		head = node
	}

	return head
}

func stacking(p *ListNode, stack *[]int) {
	for p != nil {
		*stack = append(*stack, p.Val)
		p = p.Next
	}
}

// 该题可以分解成几个小的部分，值得细细研究。
// 1.处理链表[7,7,10,7]中的进位信息，预期是[7,8,0,7]。参考：addTwoNumbers4第三步实现，最简洁。
// 2.两个栈相加，返回链表。[7,2,4,3]，[5,6,4]预期输出[7,8,0,7]。参考：addTwoNumbers5第三步实现，最简洁。
