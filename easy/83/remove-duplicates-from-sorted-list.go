package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next
	for curr != nil {
		if prev.Val == curr.Val {
			prev.Next = prev.Next.Next
			curr = curr.Next
		} else {
			curr = curr.Next
			prev = prev.Next
		}
	}

	return head
}
