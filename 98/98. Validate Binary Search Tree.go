package _8

import "math"

// see https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 4 ms, faster than 98.26% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.3 MB, less than 88.89% of Go online submissions for Validate Binary Search Tree.
*/
// 解法一：递归
// TODO 严格来说，该解法是有bug的，就是当root.Val本身的值等于math.MinInt64或math.MaxInt64时会返回false。
func isValidBST1(root *TreeNode) bool {
	return myIsValidBST(root, math.MinInt64, math.MaxInt64)
}

func myIsValidBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return myIsValidBST(root.Left, min, root.Val) && myIsValidBST(root.Right, root.Val, max)
}

/*
Runtime: 8 ms, faster than 53.56% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.9 MB, less than 11.11% of Go online submissions for Validate Binary Search Tree.
*/
// 解法二：中序遍历，然后检查是否有序
func isValidBST2(root *TreeNode) bool {
	var (
		values  []int
		inOrder func(node *TreeNode, v *[]int)
	)

	inOrder = func(node *TreeNode, v *[]int) {
		if node == nil {
			return
		}

		inOrder(node.Left, v)
		*v = append(*v, node.Val)
		inOrder(node.Right, v)
	}

	inOrder(root, &values)
	for i := 1; i < len(values); i++ {
		if values[i-1] >= values[i] {
			return false
		}
	}

	return true
}

// 和解法二一样，内存占用不一样，需要对比分析，为何内存占用不一样？
// 涉及对go底层slice的实现理解
func isValidBST3(root *TreeNode) bool {
	var (
		values  []int
		inOrder func(node *TreeNode)
	)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		values = append(values, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	for i := 1; i < len(values); i++ {
		if values[i-1] >= values[i] {
			return false
		}
	}

	return true
}
