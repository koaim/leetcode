package leetcode

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.
*/
func subarraySum(nums []int, k int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			result++
		}

		currSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			currSum += nums[j]
			if currSum == k {
				result++
			}
		}
	}

	return result
}
