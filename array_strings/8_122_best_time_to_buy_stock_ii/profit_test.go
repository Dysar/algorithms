package main

import "testing"

func Test_maxProfit(t *testing.T) {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if profit != 7 {
		t.Errorf("maxProfit() = %d, want 7", profit)
	}
}
