package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).
A valid BST is defined as follows:
  - The left subtree of a node contains only nodes with keys strictly less than the node's key.
  - The right subtree of a node contains only nodes with keys strictly greater than the node's key.
  - Both the left and right subtrees must also be binary search trees.
*/
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(root, left, right *TreeNode) bool {
	if root == nil {
		return true
	}

	if left != nil && root.Val <= left.Val {
		return false
	}

	if right != nil && root.Val >= right.Val {
		return false
	}

	return isValid(root.Left, left, root) && isValid(root.Right, root, right)
}
