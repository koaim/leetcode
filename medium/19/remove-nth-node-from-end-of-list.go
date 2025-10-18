package leetcode

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
