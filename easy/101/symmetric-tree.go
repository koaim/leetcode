package leetcode

/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricNodes(root.Left, root.Right)
}

func isSymmetricNodes(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil && r != nil || l != nil && r == nil || l.Val != r.Val {
		return false
	}

	return isSymmetricNodes(l.Left, r.Right) && isSymmetricNodes(l.Right, r.Left)
}
