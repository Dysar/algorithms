package main

import "testing"

func Test_maxProfit(t *testing.T) {
	p := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if p != 5 {
		t.Errorf("maxProfit = %d, want 5", p)
	}

}
