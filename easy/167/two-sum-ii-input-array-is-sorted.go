package leetcode

/*
Given an array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.
Return the indices of the two numbers (1-indexed) as an integer array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
*/

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)

	l := 0
	r := len(numbers) - 1

	for {
		sum := numbers[l] + numbers[r]
		if sum == target {
			res[0] = l + 1
			res[1] = r + 1
			break
		} else if sum > target {
			r--
		} else {
			l++
		}
	}

	return res
}
