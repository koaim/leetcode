package leetcode

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. 
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
*/
func removeDuplicates(nums []int) int {
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }

    return len(nums)
}