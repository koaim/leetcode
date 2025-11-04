package leetcode

import "math"

/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for i, v := range nums {
		prev, ok := m[v]
		if ok && math.Abs(float64(i-prev)) <= float64(k) {
			return true
		}
		m[v] = i
	}

	return false
}
