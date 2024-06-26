package main

import "math"

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array
// such that nums[i] == nums[j] and abs(i - j) <= k.
func containsNearbyDuplicate(nums []int, k int) bool {
	// TC: O(n), SC: O(n)

	// we map value to the index
	visited := map[int]int{} // 1,0,1,1
	for i, n := range nums {
		if index, ok := visited[n]; ok {
			if math.Abs(float64(i-index)) <= float64(k) {
				return true
			}
		}
		visited[n] = i
	}

	return false
}
