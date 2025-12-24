package leetcode

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.
*/
func findMaxConsecutiveOnes(nums []int) int {
	var left, zeroCount, maxLen int

	for right, v := range nums {
		if v == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
