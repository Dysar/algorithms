package main

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
func climbStairs(n int) int {
	// Bottom Up dynamic programming approach
	one, two := 1, 1 // if we take one step, if we take 2 steps

	for i := n - 1; i > 0; i-- {
		temp := one
		one = one + two
		two = temp
	}
	return one
}
