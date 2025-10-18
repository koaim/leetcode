package leetcode

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/
func productExceptSelf(nums []int) []int {
	zeroCount := 0
	product := 1

	for _, v := range nums {
		if v == 0 {
			zeroCount++
		} else {
			product *= v
		}
	}

	if zeroCount > 1 {
		return make([]int, len(nums))
	}

	res := make([]int, len(nums))
	for i, v := range nums {
		if zeroCount == 0 {
			res[i] = product / v
		} else if v == 0 {
			res[i] = product
		} else {
			res[i] = 0
		}
	}

	return res
}
