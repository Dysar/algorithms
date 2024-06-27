package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		l := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
		if l != 4 {
			t.Fatalf("got %d, want %d", l, 4)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		l := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
		if l != 9 {
			t.Fatalf("got %d, want %d", l, 9)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		l := longestConsecutive([]int{0, 0})
		if l != 1 {
			t.Fatalf("got %d, want %d", l, 1)
		}
	})
}
