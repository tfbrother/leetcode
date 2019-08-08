package _44

// see https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2.1 MB, less than 75.00% of Go online submissions for Binary Tree Preorder Traversal.
*/
// 解法一：递归实现
func preorderTraversal1(root *TreeNode) []int {
	var ret []int

	myPreorderTraversal(root, &ret)
	return ret
}

func myPreorderTraversal(root *TreeNode, v *[]int) {
	if root != nil {
		*v = append(*v, root.Val)
		myPreorderTraversal(root.Left, v)
		myPreorderTraversal(root.Right, v)
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
*/
// 解法二：迭代实现
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	values, stack, top := make([]int, 0), make([]*TreeNode, 0), root
	stack = append(stack, root)

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append(values, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return values
}
