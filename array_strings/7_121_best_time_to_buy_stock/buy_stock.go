package main

// buy low, sell high
// input [7,1,5,3,6,4],
// output 5: max profit 6-1. Output 0 if there is no profit to be made
// explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
func maxProfit(prices []int) int {
	// we need to find a minimum element so that it's not last
	// then starting from the next element of the minimum find max element that is greater than the minimum

	// to find minimum we can save a pointer to that index and compare next elements until last-1 is reached
	// then starting from the next index we find the max element.

	l, r := 0, 1
	maxP := 0
	for r < len(prices) { // max for r will be 5, len 6
		// if profitable
		if prices[r] > prices[l] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}

	return maxP
}
