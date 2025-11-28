package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB
	lenA, lenB := 1, 1

	for currA.Next != nil {
		currA = currA.Next
		lenA++
	}
	for currB.Next != nil {
		currB = currB.Next
		lenB++
	}

	if currA != currB {
		return nil
	}

	currA, currB = headA, headB
	i := 0
	if lenB > lenA {
		rotate := lenB - lenA
		for i != rotate {
			currB = currB.Next
			i++
		}
	} else if lenA > lenB {
		rotate := lenA - lenB
		for i != rotate {
			currA = currA.Next
			i++
		}
	}

	for currA != currB {
		currA = currA.Next
		currB = currB.Next
	}

	return currA
}
