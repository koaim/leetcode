package leetcode

/*
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.
*/
func permute(nums []int) [][]int {
	result := [][]int{}
	permuteStep(nums, 0, &result)
	return result
}

func permuteStep(nums []int, step int, result *[][]int) {
	if len(nums) == step {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
	}

	for i := step; i < len(nums); i++ {
		nums[i], nums[step] = nums[step], nums[i]
		permuteStep(nums, step+1, result)
		nums[step], nums[i] = nums[i], nums[step]
	}
}
