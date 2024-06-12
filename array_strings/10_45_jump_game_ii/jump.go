package main

//	0,1,2,3,4
//
// case 1: [2,3,1,1,4]
// jump Returns the minimum number of jumps to reach last element
// The test cases are generated such that you can reach last element
func jump(nums []int) int {
	// Brute force
	// from where can we jump to last el? From 0 to index 1, from index 1 to index 4 because 3+1 = 4.

	// i is the position. nums[i] is the jump length. We can jump from 0 to nums[i] length

	numJumps := 0
	l, r := 0, 0

	// last index
	goal := len(nums) - 1

	for r < goal {
		farthest := 0

		// window from l to (r). For example from 1 to 2, size is 1
		windowSize := r + 1 - l
		// we iterate from l index to r. i := 0; i< 0; i ++
		for i := l; i < l+windowSize; i++ {
			// compare the position we can get from other jumps in this window to the current index jump
			farthest = max(farthest, i+nums[i])
		}
		// not sure why
		l = r + 1
		// set the next window up to the farthest index we can reach
		r = farthest
		numJumps += 1
	}

	return numJumps
}
