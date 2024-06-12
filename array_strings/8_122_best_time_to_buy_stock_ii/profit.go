package main

// input prices = [7,1,5,3,6,4]
// output: 7 (4+3)
// Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		// if we will make profit
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
