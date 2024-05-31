package main

// case 1: [2,3,1,1,4]
// [ 3,2,1,0,4]
func canJump2(nums []int) bool {
	goal := len(nums) - 1

	for i := goal; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
