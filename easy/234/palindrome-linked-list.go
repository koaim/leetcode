package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/* Given the head of a singly linked list, return true if it is a palindrome or false otherwise. */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	arr := []int{head.Val}

	for head.Next != nil {
		arr = append(arr, head.Next.Val)
		head = head.Next
	}

	if len(arr) == 1 {
		return true
	}

	isMedian := true
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			isMedian = false
			break
		}
	}

	return isMedian
}
