package _4

// see https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

// 解法一：递归实现
// 该解法存在bug，因为使用了ret全局变量，且没有在函数内部初始化，多次调用inorderTraversal会存在变量污染。
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	inorderTraversal1(root.Left)
	ret = append(ret, root.Val)
	inorderTraversal1(root.Right)

	return ret
}

// 解法二
func inorderTraversal2(root *TreeNode) []int {
	ret = make([]int, 0)
	myInorderTraversal(root)

	return ret
}

func myInorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	myInorderTraversal(root.Left)
	ret = append(ret, root.Val)
	myInorderTraversal(root.Right)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
*/
// 解法三：迭代实现，借助栈
func inorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	values, stack, cur := make([]int, 0), make([]*TreeNode, 0), root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
			values = append(values, cur.Val)
			cur = cur.Right
		}
	}

	return values
}

// 解法三代码精简
func inorderTraversal3_1(root *TreeNode) []int {
	values, stack, cur := make([]int, 0), make([]*TreeNode, 0), root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append(values, cur.Val)
		cur = cur.Right
	}

	return values
}
