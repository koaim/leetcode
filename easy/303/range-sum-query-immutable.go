package leetcode

/*
Given an integer array nums, handle multiple queries of the following type:
  - Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

Implement the NumArray class:
  - NumArray(int[] nums) Initializes the object with the integer array nums.
  - int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	var sums []int
	var sum int

	for _, v := range nums {
		sum += v
		sums = append(sums, sum)
	}

	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	rightV := this.sum[right]
	if left == 0 {
		return rightV
	}

	leftV := this.sum[left-1]
	return rightV - leftV
}
