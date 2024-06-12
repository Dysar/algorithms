package main

import (
	"fmt"
	"slices"
)

// For example [3,2,2,3],3 is the element to remove.
// returns the number of elements which are not equal to val
func removeElement(nums []int, val int) int {
	var n int
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] != val {
				nums[n] = nums[i]
				n++
			}
		}
	}
	return n
}

// execution comments
// i = 0, n = 0; [
// i = 1, n = 1. [2,
// i = 2; n = 2  [2,2
// i = 3; n = 2  [2,2

func main() {
	nums := []int{3, 2, 2, 3}   // Input array
	val := 3                    // Value to remove
	expectedNums := []int{2, 2} // The expected answer with correct length.
	// It is sorted with no values equaling val.

	k := removeElement(nums, val) // Calls your implementation

	if k != len(expectedNums) {
		fmt.Printf("k = %d, expectedNums = %d\n", k, len(expectedNums))
		return
	}
	slices.Sort(nums) // Sort the elements of nums

	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			fmt.Printf("nums[%d] = %d\n", i, nums[i])
		}
	}

}
