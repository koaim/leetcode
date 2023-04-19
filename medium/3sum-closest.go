package leetcode

import "sort"

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.

You may assume that each input would have exactly one solution.
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			deltaSum := abs(target - sum)
			deltaResult := abs(target - result)

			if deltaSum < deltaResult {
				result = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}
