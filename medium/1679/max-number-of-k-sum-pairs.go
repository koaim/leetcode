package leetcode

import "slices"

/*
You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.
*/
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)

	left, right := 0, len(nums)-1
	res := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			left++
			right--
			res++
		} else if sum > k {
			right--
		} else {
			left++
		}
	}

	return res
}
