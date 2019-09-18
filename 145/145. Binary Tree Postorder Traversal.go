package _45

// see https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 75.00% of Go online submissions for Binary Tree Postorder Traversal.
*/
// 解法一：递归实现
func postorderTraversal1(root *TreeNode) []int {
	var ret []int

	myPostOrderTraversal(root, &ret)
	return ret
}

func myPostOrderTraversal(root *TreeNode, v *[]int) {
	if root != nil {
		myPostOrderTraversal(root.Left, v)
		myPostOrderTraversal(root.Right, v)
		*v = append(*v, root.Val)
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
*/
// 解法二：迭代实现
// 采用根右左遍历，然后倒序输出就是左右根后序遍历了
func postorderTraversal2(root *TreeNode) []int {
	ret := preorder(root)
	reverseSlice(ret)
	return ret
}

// 根右左遍历
func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	values, stack, top := make([]int, 0), make([]*TreeNode, 0), root
	stack = append(stack, root)

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		values = append(values, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	return values
}

// 反转slice
func reverseSlice(v []int) {
	l, r := 0, len(v)-1
	for l < r {
		v[l], v[r] = v[r], v[l]
		l++
		r--
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
*/
// 解法三：迭代实现
// 采用根右左遍历，然后倒序输出就是左右根后序遍历了
func postorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	values, stack, cur := make([]int, 0), make([]*TreeNode, 0), root
	stack = append(stack, cur)
	stack = append(stack, cur)

	for len(stack) > 0 {
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if len(stack) > 0 && cur == stack[len(stack)-1] {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				stack = append(stack, cur.Left)
			}
		} else {
			values = append(values, cur.Val)
		}
	}

	return values
}

// 解法四：迭代实现，模拟系统栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack, ret := make([]*commond, 1), make([]int, 0)
	stack[0] = &commond{s: "go", node: root}
	var top *commond

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if top.s == "go" { // 访问
			stack = append(stack, &commond{s: "print", node: top.node})
			if top.node.Right != nil {
				stack = append(stack, &commond{s: "go", node: top.node.Right})
			}
			if top.node.Left != nil {
				stack = append(stack, &commond{s: "go", node: top.node.Left})
			}
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
