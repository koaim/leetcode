package leetcode

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	m := map[int]bool{}

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	var r int
	for k := range m {
		r = k
	}

	return r
}
