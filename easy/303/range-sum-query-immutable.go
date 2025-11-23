package leetcode

type NumArray struct {
	nums []int
}

/*
Given an integer array nums, handle multiple queries of the following type:
  - Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

Implement the NumArray class:
  - NumArray(int[] nums) Initializes the object with the integer array nums.
  - int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/
func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}

	return sum
}
