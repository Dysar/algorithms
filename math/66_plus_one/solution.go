package main

// sample input digits = [1,2,3]
// sample output [1,2,4]
// explanation: 123 + 1 = 124
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	// newDigits is going to be like 10000000.
	return newDigits

	// TC: O(n), SC: O(1)
}
