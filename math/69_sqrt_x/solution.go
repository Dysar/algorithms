package main

// x > 0, return sqrt
func mySqrt(x int) int {
	//sqrt of 0 is 0, of 1 is 1.
	if x < 1 {
		return 0
	}

	if x < 4 {
		return 1
	}

	start, end, mid := 1, x, -1
	for start <= end {
		// Calculate the middle point of the set
		mid = start + (end-start)/2

		// If the square of the middle value is greater than x, move the "end" to the left subset to traverse it
		if mid*mid > x {
			end = mid - 1
		} else if mid*mid == x {
			// If the square of the middle value is equal to x, we found the square root.
			return mid
		} else {
			// If the square of the middle value is less than x, move the "start" to the right subset
			start = mid + 1
		}
	}

	// The loop ends when "start" becomes greater than "end", and "end" is the integer value of the square root.
	return end
}
