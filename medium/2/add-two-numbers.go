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
	r := &ListNode{0, nil}
	c := r
	s := 0

	for {
		if l1 == nil && l2 == nil {
			if s != 0 {
				c.Next = &ListNode{Val: s}
			}
			break
		}

		var v1, v2 int
		if l1 != nil {
			v1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			v2, l2 = l2.Val, l2.Next
		}

		v := v1 + v2 + s
		if v >= 10 {
			v, s = v-10, 1
		} else {
			s = 0
		}

		c.Next = &ListNode{Val: v}
		c = c.Next
	}

	return r.Next
}
