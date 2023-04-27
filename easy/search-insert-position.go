package leetcode

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		center := left + (right-left)/2

		if nums[center] > target {
			right = center - 1
		} else if nums[center] < target {
			left = center + 1
		} else {
			return center
		}
	}

	return left
}
