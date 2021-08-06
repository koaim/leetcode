package leetcode

/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	r := -1
	k := len(nums) / 2
	v := nums[k]

	if v == target {
		r = k
	} else if v < target {
		r = search(nums[k+1:], target)
		if r != -1 {
			r += k + 1
		}
	} else {
		r = search(nums[:k], target)
	}

	return r
}
