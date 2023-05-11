package leetcode

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/
func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}

	return false
}
