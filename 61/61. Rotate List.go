package _1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：暴力解法
// 优化次数，获取链表的个数，然后与k取余数
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur, count := head, 0
	for cur != nil {
		cur, count = cur.Next, count+1
	}

	k %= count // TODO 第一版就是没有这个逻辑，导致超时了。
	for i := 0; i < k; i++ {
		head = rotateRightOne(head)
	}

	return head
}

// 右移动一位，其实就是把尾部的节点放到头部来
func rotateRightOne(head *ListNode) *ListNode {
	prev, cur, next := head, head, head.Next

	for next != nil {
		prev, cur, next = cur, next, next.Next
	}

	// next == nil, 此时cur指向的就是末尾的元素

	prev.Next, cur.Next = nil, head
	return cur
}
