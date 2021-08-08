package leetcode

/*
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i := len(nums) / 2
	v := nums[i]
	var r int

	if v == target {
		r = i
	} else if nums[0] < v {
		if nums[0] <= target && v >= target {
			r = search(nums[:i], target)
		} else {
			if r = search(nums[i+1:], target); r != -1 {
				r += i + 1
			}
		}
	} else {
		if v <= target && nums[len(nums)-1] >= target {
			if r = search(nums[i+1:], target); r != -1 {
				r += i + 1
			}
		} else {
			r = search(nums[:i], target)
		}
	}

	return r
}
