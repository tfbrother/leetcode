package _3__Merge_k_Sorted_Lists

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 948 ms, faster than 5.11% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 6 MB, less than 21.74% of Go online submissions for Merge k Sorted Lists.
*/
// 解法一：K路归并排序。
// 直观想法，就是采用K路归并排序，只是对象变成链表而已。
// 时间复杂度O(NK)，空间复杂度O(K)
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		head, cur, prev     *ListNode
		minVal, minIndex, k int
		listsMap            map[int]*ListNode
	)

	head, listsMap = &ListNode{}, make(map[int]*ListNode)
	prev = head

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			listsMap[i] = lists[i]
		}
	}

	for len(listsMap) > 0 {
		minVal = math.MaxInt32
		// 查找最小值
		for k, cur = range listsMap {
			if minVal > cur.Val {
				minVal = cur.Val
				minIndex = k
			}
		}

		prev.Next = listsMap[minIndex]
		prev = prev.Next
		if prev.Next == nil {
			delete(listsMap, minIndex)
		} else {
			listsMap[minIndex] = prev.Next
		}
	}

	return head.Next
}

/*
Runtime: 8 ms, faster than 97.12% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 5.3 MB, less than 100.00% of Go online submissions for Merge k Sorted Lists.
*/
// 解法二：真正的K路归并排序
// 解法一其实不算K路归并排序，因为并没有实现归并，其实是逐一比较而已。真正的归并算法需要两两合并，直到只有一个。
// 采用真正的归并排序，时间复杂度O(NlogK)，空间复杂度为O(1)
func mergeKLists2(lists []*ListNode) *ListNode {
	// TODO 如何归并？递归
	// 要把k个分成两部分，递归分解，直到只剩下两部分。
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeLists(lists[0], lists[1])
	}

	middle := len(lists) >> 1
	left, right := mergeKLists2(lists[:middle]), mergeKLists2(lists[middle:])
	return mergeLists(left, right)
}

// 合并两个有序链表
func mergeLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			prev, l1 = prev.Next, l1.Next
		} else {
			prev.Next = l2
			prev, l2 = prev.Next, l2.Next
		}
	}

	if l1 == nil {
		l1 = l2
	}

	for l1 != nil {
		prev.Next = l1
		prev, l1 = prev.Next, l1.Next
	}

	return head.Next
}

// 要深究解法一和解法二为何时间复杂度不一样。K路归并排序效率是没有2路归并排序效率高的。
// 归并排序还是没有深入理解清楚啊。包括归并排序的时间复杂度分析也没搞明白。
