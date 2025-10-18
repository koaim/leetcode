package leetcode

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}
