package leetcode

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
*/

func missingNumber(nums []int) int {
	var actual int
	for _, v := range nums {
		actual += v
	}

	var want int
	for i := 0; i < len(nums)+1; i++ {
		want += i
	}

	return want - actual
}
