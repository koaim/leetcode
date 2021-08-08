package leetcode

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	r := make([]int, 2)
	m := map[int]int{}

	for i, v := range nums {
		number := target - v
		vv, ok := m[number]
		if ok {
			r[0] = i
			r[1] = vv
			break
		} else {
			m[v] = i
		}
	}

	return r
}
