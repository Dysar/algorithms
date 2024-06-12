package main

import "testing"

func Test_hindex(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		citations := []int{3, 0, 6, 1, 5}
		res := hIndex(citations)
		if res != 3 {
			t.Errorf("hIndex(%v) got %d, want 3", citations, res)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		citations := []int{1, 3, 1}
		res := hIndex(citations)
		if res != 1 {
			t.Errorf("hIndex(%v) got %d, want 1", citations, res)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		citations := []int{100}
		res := hIndex(citations)
		if res != 1 {
			t.Errorf("hIndex(%v) got %d, want 1", citations, res)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		citations := []int{11, 15}
		res := hIndex(citations)
		if res != 2 {
			t.Errorf("hIndex(%v) got %d, want 2", citations, res)
		}
	})
}
