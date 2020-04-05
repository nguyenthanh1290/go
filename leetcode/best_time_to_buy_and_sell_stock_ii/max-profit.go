package best_time_to_buy_and_sell_stock_ii

// Time complexity : O(n)
// Space complexity: O(1)
func maxProfit(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
