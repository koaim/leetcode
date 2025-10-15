package leetcode

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
