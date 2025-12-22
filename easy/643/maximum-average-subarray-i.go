package leetcode

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
Any answer with a calculation error less than 10-5 will be accepted.
*/
func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
