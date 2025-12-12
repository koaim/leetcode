package leetcode

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	curr := len(nums) - 1

	for curr >= 0 {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if rightSquare > leftSquare {
			res[curr] = rightSquare
			right--
		} else {
			res[curr] = leftSquare
			left++
		}

		curr--
	}

	return res
}
