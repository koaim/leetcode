package leetcode

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	t := head
	for i := 0; i < n; i++ {
		t = t.Next
	}

	if t == nil {
		return head.Next
	}

	cur := head
	var prev *ListNode

	for {
		t = t.Next
		prev = cur
		cur = cur.Next

		if t == nil {
			break
		}
	}

	prev.Next = cur.Next

	return head
}
