package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i := len(nums) / 2
	v := nums[i]
	var r int

	if v == target {
		r = i
	} else if nums[0] < v {
		if nums[0] <= target && v >= target {
			r = search(nums[:i], target)
		} else {
			if r = search(nums[i+1:], target); r != -1 {
				r += i + 1
			}
		}
	} else {
		if v <= target && nums[len(nums)-1] >= target {
			if r = search(nums[i+1:], target); r != -1 {
				r += i + 1
			}
		} else {
			r = search(nums[:i], target)
		}
	}

	return r
}
