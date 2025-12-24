package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].
*/
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum int
	if root.Val >= low && root.Val <= high {
		sum = root.Val
	}

	sum += rangeSumBST(root.Right, low, high)
	sum += rangeSumBST(root.Left, low, high)

	return sum
}
