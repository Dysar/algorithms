package main

import "testing"

func Test_mySqrt(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		num := mySqrt(4)
		if num != 2 {
			t.Errorf("expected: %d, got: %d", 4, num)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		num := mySqrt(17)
		if num != 4 {
			t.Errorf("expected: %d, got: %d", 4, num)
		}
	})
}
