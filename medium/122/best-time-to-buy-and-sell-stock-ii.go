package leetcode

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
However, you can sell and buy the stock multiple times on the same day, ensuring you never hold more than one share of the stock.

Find and return the maximum profit you can achieve.
*/
func maxProfit(prices []int) int {
	buy, sell := -1, -1
	canBuy := true
	wasSell := false

	var sum, lastDeal int
	for _, v := range prices {
		if wasSell && v > sell {
			sum -= lastDeal
			sell = v
			lastDeal = sell - buy
			sum += lastDeal
			canBuy = true
			wasSell = true
		} else if canBuy {
			buy = v
			canBuy = false
			wasSell = false
		} else if !canBuy && v < buy {
			buy = v
		} else if !canBuy && v > buy {
			sell = v
			lastDeal = sell - buy
			sum += lastDeal
			canBuy = true
			wasSell = true
		}
	}

	return sum
}
