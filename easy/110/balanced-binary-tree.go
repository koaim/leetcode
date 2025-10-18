package leetcode

/*
Given a binary tree, determine if it is height-balanced.
*/

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := check(root)
	return ok
}

func check(n *TreeNode) (float64, bool) {
	if n == nil {
		return 0, true
	}

	l, ok := check(n.Left)
	if !ok {
		return 0, false
	}

	r, ok := check(n.Right)
	if !ok {
		return 0, false
	}

	if diff := math.Abs(l - r); diff > 1 {
		return 0, false
	}

	return math.Max(l, r) + 1, true
}
