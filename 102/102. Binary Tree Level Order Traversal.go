package _02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	level int       // 层数
	node  *TreeNode // 结点
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
*/
// 解法一：BFS
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	list, ret := make([]*item, 1), make([][]int, 0)
	list[0] = &item{level: 0, node: root}

	var top *item
	for len(list) > 0 {
		top, list = list[0], list[1:]
		if len(ret) == top.level {
			ret = append(ret, []int{top.node.Val})
		} else {
			ret[top.level] = append(ret[top.level], top.node.Val)
		}
		if top.node.Left != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Left})
		}
		if top.node.Right != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Right})
		}
	}

	return ret
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 6.1 MB, less than 50.00% of Go online submissions for Binary Tree Level Order Traversal.
*/
// 解法二：DFS
func levelOrder2(root *TreeNode) [][]int {
	ret := make([][]int, 0)

	DFS(root, 0, &ret)
	return ret
}

func DFS(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if len(*ret) == level {
		*ret = append(*ret, []int{})
	}

	(*ret)[level] = append((*ret)[level], root.Val)
	DFS(root.Left, level+1, ret)
	DFS(root.Right, level+1, ret)
}
