package leetcode

/*
We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.
*/
func findLHS(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	var res int
	for _, v := range nums {
		countGreater, ok := m[v+1]
		if ok {
			count := m[v] + countGreater
			if count > res {
				res = count
			}
		}
	}

	return res
}
