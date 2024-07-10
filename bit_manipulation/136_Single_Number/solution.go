package main

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
