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

// 解法三：迭代实现，模拟系统栈
// 该解法很容易的改造成中序后序的迭代实现。
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack, ret := make([]*commond, 1), make([]int, 0)
	stack[0] = &commond{s: "go", node: root}
	var top *commond

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if top.s == "go" { // 访问
			if top.node.Right != nil {
				stack = append(stack, &commond{s: "go", node: top.node.Right})
			}
			if top.node.Left != nil {
				stack = append(stack, &commond{s: "go", node: top.node.Left})
			}
			stack = append(stack, &commond{s: "print", node: top.node})
		} else {
			ret = append(ret, top.node.Val)
		}
	}

	return ret
}

type commond struct {
	s    string
	node *TreeNode
}
