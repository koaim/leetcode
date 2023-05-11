package leetcode

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	pos := []int{}
	neg := []int{}

	for _, v := range nums {
		if v >= 0 {
			pos = append(pos, v*v)
		} else {
			neg = append([]int{v * v}, neg...)
		}
	}

	if len(pos) == 0 {
		return neg
	}

	if len(neg) == 0 {
		return pos
	}

	res := []int{}
	i, j := 0, 0

	for len(res) != len(nums) {
		if i == len(pos) {
			res = append(res, neg[j])
			j++
		} else if j == len(neg) {
			res = append(res, pos[i])
			i++
		} else {
			if pos[i] < neg[j] {
				res = append(res, pos[i])
				i++
			} else {
				res = append(res, neg[j])
				j++
			}
		}
	}

	return res
}
