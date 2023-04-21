package _go

// Needs improvement, 202/211 test cases passed, "time limit exceeded"
func maxProfit(prices []int) int {
	var maxProfitBuy, maxProfitSell, maxProfit int
	buy := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		}
		for j := i + 1; j < len(prices); j++ {
			possibleMaxProfit := prices[j] - buy
			if possibleMaxProfit > maxProfit {
				maxProfit = possibleMaxProfit
				maxProfitBuy = buy
				maxProfitSell = prices[j]
			}
		}
	}
	return maxProfitSell - maxProfitBuy
}
