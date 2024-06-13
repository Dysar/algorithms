package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	div := 1

	// here we set the divider to the len of x. x > 10 true, x> 100 true etc
	for x >= 10*div {
		div *= 10
	}

	for x > 0 {
		rightDigit := x % 10
		leftDigit := x / div

		// compare left and right digits 121
		if leftDigit != rightDigit {
			return false
		}

		// chop off left and right digit
		x = (x % div) / 10

		// chop off 2 digits that we've just compared
		div = div / 100
	}
	return true
}
