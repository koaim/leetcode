package main

/*
Given the root of a binary tree and an integer targetSum.
Return all root-to-leaf paths where the sum of the node values in the path equals targetSum.
Each path should be returned as a list of the node values, not node references.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	dfs(root, []int{}, &res, 0, targetSum)

	return res
}

func dfs(n *TreeNode, path []int, res *[][]int, sum, targetSum int) {
	if n == nil {
		return
	}

	path = append(path, n.Val)
	sum += n.Val

	if sum == targetSum && n.Left == nil && n.Right == nil {
		*res = append(*res, path)
	}

	p := append(make([]int, 0, len(path)), path...)

	dfs(n.Left, p, res, sum, targetSum)
	dfs(n.Right, p, res, sum, targetSum)
}
