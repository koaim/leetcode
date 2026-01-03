package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
*/
func rightSideView(root *TreeNode) []int {
	var elements [][]int

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(elements) == level {
			elements = append(elements, []int{})
		}

		elements[level] = append(elements[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	var result []int
	for _, v := range elements {
		last := v[len(v)-1]
		result = append(result, last)
	}

	return result
}
