package leetcode

/*
Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
*/
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1

	if nums[left] < nums[right] {
		return nums[left]
	}

	for right >= left {
		mid := left + (right-left)/2

		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		} else if mid < len(nums)-1 && nums[mid+1] < nums[mid] {
			return nums[mid+1]
		} else if nums[mid] > nums[0] {
			left++
		} else {
			right--
		}
	}

	return -1
}
