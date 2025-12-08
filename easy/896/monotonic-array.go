package leetcode

/*
An array is monotonic if it is either monotone increasing or monotone decreasing.
An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].

Given an integer array nums, return true if the given array is monotonic, or false otherwise.
*/
func isMonotonic(nums []int) bool {
	var up, down bool
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			down = true
		} else if nums[i] > nums[i-1] {
			up = true
		} else {
			continue
		}

		if up && down {
			return false
		}
	}

	return true
}
