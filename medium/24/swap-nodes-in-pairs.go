package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given a linked list, swap every two adjacent nodes and return its head.
You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
*/
func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	curr := res

	for curr.Next != nil && curr.Next.Next != nil {
		next2 := curr.Next.Next
		next := curr.Next

		next.Next = next2.Next
		next2.Next = next

		curr.Next = next2
		curr = next
	}

	return res.Next
}
