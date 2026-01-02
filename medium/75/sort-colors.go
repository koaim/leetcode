package leetcode

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/
func sortColors(nums []int) {
	var zeroCount, oneCount, twoCount int

	for _, v := range nums {
		switch v {
		case 0:
			zeroCount++
		case 1:
			oneCount++
		default:
			twoCount++
		}
	}

	for i := range nums {
		if zeroCount > 0 {
			nums[i] = 0
			zeroCount--
		} else if oneCount > 0 {
			nums[i] = 1
			oneCount--
		} else {
			nums[i] = 2
			twoCount--
		}
	}
}
