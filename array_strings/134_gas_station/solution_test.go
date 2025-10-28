package main

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		gas := []int{1, 2, 3, 4, 5}
		cost := []int{3, 4, 5, 1, 2}
		want := 3
		got := canCompleteCircuit(gas, cost)
		if got != want {
			t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", gas, cost, got, want)
		}
	})
}
