package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	var helper func(node *TreeNode, level int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)

		helper(node.Left, level+1)
		helper(node.Right, level+1)
	}

	helper(root, 0)
	return res
}
