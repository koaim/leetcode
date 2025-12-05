package leetcode

func countPartitions(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	var count, currSum int
	for _, v := range nums {
		currSum += v
		otherSum := sum - currSum
		if (currSum-otherSum)%2 == 0 {
			count++
		}
	}

	if count > 1 {
		count--
	}

	return count
}
