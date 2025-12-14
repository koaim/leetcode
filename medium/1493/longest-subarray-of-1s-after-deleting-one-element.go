package leetcode

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.
*/
func longestSubarray(nums []int) int {
	var prevZero, count, maxCount int
	var wasZero bool

	for i, v := range nums {
		if v == 1 {
			count++
		} else if v == 0 && !wasZero {
			prevZero = i
			wasZero = true
		} else {
			maxCount = max(count, maxCount)
			count = i - prevZero - 1
			prevZero = i
		}
	}

	maxCount = max(count, maxCount)
	if !wasZero {
		return len(nums) - 1
	}

	return maxCount
}
