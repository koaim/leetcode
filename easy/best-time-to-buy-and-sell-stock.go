package leetcode

func maxProfit(prices []int) int {
	minPrev := prices[0]
	result := 0

	for i := 1; i < len(prices); i++ {
		res := prices[i] - minPrev
		if res > result {
			result = res
		}

		minPrev = min(minPrev, prices[i])
	}

	return result
}
