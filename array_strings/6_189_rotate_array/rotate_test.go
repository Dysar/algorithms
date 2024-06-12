package main

import "testing"

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{4, 5, 1, 2, 3}
	rotate(nums, 2)

	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("rotate got %v, expected %v at %d", nums[i], expected[i], i)
		}
	}
}
