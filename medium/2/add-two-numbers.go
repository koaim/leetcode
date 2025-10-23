package leetcode

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	v1, v2, v3 := l1, l2, l3
	var delta int

	for v1 != nil || v2 != nil {
		var val1, val2 int

		if v1 != nil {
			val1 = v1.Val
		}

		if v2 != nil {
			val2 = v2.Val
		}

		val3 := val1 + val2 + delta
		if val3 > 9 {
			val3 = val3 - 10
			delta = 1
		} else {
			delta = 0
		}

		v3.Next = &ListNode{Val: val3}
		v3 = v3.Next

		if v1 != nil {
			v1 = v1.Next
		}
		if v2 != nil {
			v2 = v2.Next
		}
	}

	if delta == 1 {
		v3.Next = &ListNode{Val: 1}
	}

	return l3.Next
}
