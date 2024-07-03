package main

import "testing"

func Test_climbStairs(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := climbStairs(2)
		if output != 2 {
			t.Errorf("expected: %d, got: %d", 2, output)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := climbStairs(3)
		if output != 3 {
			t.Errorf("expected: %d, got: %d", 3, output)
		}
	})
}
