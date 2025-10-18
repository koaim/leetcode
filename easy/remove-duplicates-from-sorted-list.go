package leetcode

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
