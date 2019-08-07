package _01

// see https://leetcode-cn.com/problems/symmetric-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Symmetric Tree.
*/
// 解法一：参考[100].相同的树的解法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return myIsSymmetric(root.Left, root.Right)
}

func myIsSymmetric(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 != nil {
		if root1.Val != root2.Val {
			return false
		}

		return myIsSymmetric(root1.Left, root2.Right) && myIsSymmetric(root1.Right, root2.Left)
	} else {
		return false
	}
}
