package letcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var level int

	var helper func(node *TreeNode, level int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level == len(res) {
			res = append(res, []int{})
		}

		if level%2 == 0 {
			res[level] = append(res[level], node.Val)
		} else {
			res[level] = append([]int{node.Val}, res[level]...)
		}

		helper(node.Left, level+1)
		helper(node.Right, level+1)
	}

	helper(root, level)

	return res
}
