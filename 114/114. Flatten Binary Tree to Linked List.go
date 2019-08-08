package _14

// see https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第0版本思路：题目实际考察的就是二叉树的前序遍历
func flatten0(root *TreeNode) {
	prevOrder(root)
}

func prevOrder(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left, right := root.Left, root.Right
	root.Left, root.Right = nil, left
	if left != nil {
		left = prevOrder(root.Left)
	}

	left.Right = prevOrder(right)

	return root
}

/*
Runtime: 4 ms, faster than 81.35% of Go online submissions for Flatten Binary Tree to Linked List.
Memory Usage: 6.8 MB, less than 100.00% of Go online submissions for Flatten Binary Tree to Linked List.
*/
// 解法一：递归实现，后序遍历
// 参考题解实现
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	// 后序遍历
	if root.Left != nil {
		pre := root.Left       // 令 pre 指向左子树
		for pre.Right != nil { // 找到左子树中的最右节点
			pre = pre.Right
		}

		pre.Right = root.Right // 令左子树中的最右节点的右子树 指向 根节点的右子树
		root.Right = root.Left // 令根节点的右子树指向根节点的左子树
		root.Left = nil        // 置空根节点的左子树
	}

	root = root.Right // 令当前节点指向下一个节点
}
