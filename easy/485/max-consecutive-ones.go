package leetcode

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.
*/
func findMaxConsecutiveOnes(nums []int) int {
	var length, maxLength int

	for _, v := range nums {
		if v == 1 {
			length++
		} else {
			maxLength = max(maxLength, length)
			length = 0
		}
	}

	maxLength = max(maxLength, length)

	return maxLength
}
