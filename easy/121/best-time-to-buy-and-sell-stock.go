package leetcode

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

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
