package leetcode

/*
You are given an array of integers nums.

Return the length of the longest subarray of nums which is either strictly increasing or strictly decreasing.
*/
func longestMonotonicSubarray(nums []int) int {
	maxIncLen, maxDecLen := 1, 1
	currIncLen, currDecLen := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currIncLen++
			currDecLen = 1
		} else if nums[i] < nums[i-1] {
			currDecLen++
			currIncLen = 1
		} else {
			currDecLen = 1
			currIncLen = 1
		}

		maxIncLen = max(maxIncLen, currIncLen)
		maxDecLen = max(maxDecLen, currDecLen)
	}

	return max(maxIncLen, maxDecLen)
}
