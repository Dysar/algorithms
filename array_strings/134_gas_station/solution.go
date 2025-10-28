package main

// Greedy approach Time Complexity: O(n) Space Complexity: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	// if total gas is less than total cost, it's not possible to complete the circuit
	if totalGas < totalCost {
		return -1
	}

	// get diffs
	total := 0
	start := 0
	// example gas [1,2,3,4,5] cost [3,4,5,1,2]
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]

		// 1-3 = -2 -> start = 1
		// 2-4 = -2 -> start = 2
		// 3-5 = -2 -> start = 3
		// 4-1 = 3 -> start = 3
		// 5-2 = 3 -> start = 3
		if total < 0 {
			// Greedy idea: if the partial sum dips below 0 at i, any start before i fails; move start to i+1.
			total = 0
			start = i + 1
		}
	}

	return start
}
