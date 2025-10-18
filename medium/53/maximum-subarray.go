package leetcode

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
