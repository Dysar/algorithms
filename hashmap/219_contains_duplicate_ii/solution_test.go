package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		contains := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
		if contains != true {
			t.Fatalf("expected: %v, got: %v", true, contains)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		contains := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
		if contains != true {
			t.Fatalf("expected: %v, got: %v", true, contains)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		contains := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
		if contains != false {
			t.Fatalf("expected: %v, got: %v", false, contains)
		}
	})
}
