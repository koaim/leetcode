package leetcode

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.
*/
func findDuplicate(nums []int) int {
	fast, slow := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
