package _99

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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Binary Tree Right Side View.
*/
// 思路一：BFS 解法和102完全一样，只是在返回的时只返回ret子项的最后一个元素
func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := make([]item, 1)
	ret := make([][]int, 1)
	//queue = append(queue, item{0, root})
	queue[0] = item{0, root}

	for len(queue) > 0 {
		q := queue[0]
		node := q.node
		if len(ret) == q.level {
			ret = append(ret, []int{node.Val})
		} else {
			ret[q.level] = append(ret[q.level], node.Val)
		}

		if node.Left != nil {
			queue = append(queue, item{q.level + 1, node.Left})
		}
		if node.Right != nil {
			queue = append(queue, item{q.level + 1, node.Right})
		}

		queue = queue[1:]
	}

	arr := make([]int, len(ret))
	for i := 0; i < len(ret); i++ {
		arr[i] = ret[i][len(ret[i])-1]
	}

	return arr
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Binary Tree Right Side View.
*/
// 解法二：BFS 层序遍历，只把每层最右边的数据放入结果集
// 比解法一更优
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue, ret := make([]item, 1), make([]int, 0)
	queue[0] = item{0, root}

	for len(queue) > 0 {
		width := len(queue)
		for i := 0; i < width; i++ {
			q := queue[i]
			if q.node.Left != nil {
				queue = append(queue, item{q.level + 1, q.node.Left})
			}
			if q.node.Right != nil {
				queue = append(queue, item{q.level + 1, q.node.Right})
			}
		}
		ret = append(ret, queue[width-1].node.Val)
		queue = queue[width:]
	}

	return ret
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Binary Tree Right Side View.
*/
// 解法三：DFS
func rightSideView3(root *TreeNode) []int {
	ret := make([]int, 0)
	DFS(root, 0, &ret)
	return ret
}

func DFS(root *TreeNode, level int, ret *[]int) {
	if root == nil {
		return
	}

	if level == len(*ret) {
		*ret = append(*ret, root.Val)
	}

	DFS(root.Right, level+1, ret)
	DFS(root.Left, level+1, ret)
}
