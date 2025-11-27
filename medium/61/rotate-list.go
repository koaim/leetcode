package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given the head of a linked list, rotate the list to the right by k places.
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length := 1
	top := head

	for top.Next != nil {
		top = top.Next
		length++
	}

	if k > length {
		k = k % length
	}

	if k == 0 || k == length {
		return head
	}

	i := 1
	curr := head
	res := curr.Next

	for i != length-k {
		curr = curr.Next
		res = curr.Next
		i++
	}

	curr.Next = nil
	top.Next = head

	return res
}
