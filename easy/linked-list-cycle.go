package leetcode

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := map[*ListNode]bool{}
	hasCycle := false

	for {
		if _, ok := m[head]; ok {
			hasCycle = true
			break
		}

		if head.Next == nil {
			break
		}

		m[head] = true
		head = head.Next
	}

	return hasCycle
}
