package leetcode

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
func rotate(nums []int, k int) {
	i := 0
	for i < k {
		shiftOne(nums)
		i++
	}
}

func shiftOne(a []int) {
	if len(a) <= 1 {
		return
	}

	last := a[len(a)-1]
	for i := len(a) - 1; i > 0; i-- {
		a[i] = a[i-1]
	}

	a[0] = last
}
