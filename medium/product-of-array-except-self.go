package leetcode

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	product := 1

	for i := 0; i < len(nums); i++ {
		product = product * nums[i]
		prefix[i] = product
	}

	postfix := make([]int, len(nums))
	product = 1
	for i := 0; i < len(nums); i++ {
		product = product * nums[len(nums)-1-i]
		postfix[len(nums)-1-i] = product
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pre := 1
		if i > 0 {
			pre = prefix[i-1]
		}

		post := 1
		if i != len(nums)-1 {
			post = postfix[i+1]
		}

		result[i] = pre * post
	}

	return result
}
