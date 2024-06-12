package main

import (
	"slices"
	"testing"
)

func Test_remove(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}      // Input array
	expectedNums := []int{1, 1, 2, 2, 3} // The expected answer with correct length.
	// It is sorted with no values equaling val.

	k := removeDuplicatesV2(nums) // Calls your implementation

	if k != len(expectedNums) {
		t.Errorf("k = %d, expectedNums = %d\n", k, len(expectedNums))
		return
	}
	slices.Sort(nums) // Sort the elements of nums
	t.Log(nums)

	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Errorf("unexpected element nums[%d] = %d\n", i, nums[i])
		}
	}
}
