package main

//	0,1,2,3,4
//
// case 1: [2,3,1,1,4]
func jump(nums []int) bool {
	// Return the minimum number of jumps to reach last element
	// The test cases are generated such that you can reach last element

	// Brute force
	// from where can we jump to last el? From 1 at index 3, from 3 at index 1 because 3+1 = 4.
	// from where can we jump to 3? from index 0. 2 jumps

	goal := len(nums) - 1

	for i := goal; i >= 0; i-- {
		//
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
