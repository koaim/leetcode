package leetcode

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. 
Then return the number of elements in nums which are not equal to val.
*/
func removeElement(nums []int, val int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }

    return len(nums)
}