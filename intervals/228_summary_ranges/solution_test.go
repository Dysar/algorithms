package main

import "testing"

func Test_summaryRanges(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := summaryRanges([]int{0, 1, 2, 4, 5, 7})
		if len(output) != 3 {
			t.Fatalf("Test failed: %v", output)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
		if len(output) != 4 {
			t.Fatalf("Test failed: %v", output)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		output := summaryRanges([]int{0, 1, 2})
		if len(output) != 1 {
			t.Fatalf("Test failed: %v", output)
		}
	})
}
