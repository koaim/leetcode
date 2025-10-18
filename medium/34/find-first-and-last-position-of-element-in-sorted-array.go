package leetcode

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		center := (left + right) / 2

		if nums[center] < target {
			left = center + 1
		} else {
			right = center
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		center := (left+right)/2 + 1

		if nums[center] <= target {
			left = center
		} else if nums[center] > target {
			right = center - 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}
