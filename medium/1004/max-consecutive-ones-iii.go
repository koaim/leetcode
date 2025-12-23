package leetcode

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
*/
func longestOnes(nums []int, k int) int {
	var left, zeroCount, maxLength int

	for right, v := range nums {
		if v == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
