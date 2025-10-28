package main

import "github.com/sirupsen/logrus"

// 1 <= num <= 3999
func intToRoman(num int) string {
	roman := ""
	for num > 0 {
		logrus.Infof("num: %d", num)
		if num >= 1000 {
			roman += getDigit("M", "D", "C", num, 1000)
			logrus.Infof("roman: %s", roman)
			// if num is 3999, we want 999 remaining.
			num %= 1000
		} else if num >= 100 { // it can be 9xx,8xx,7xx,6xx,5xx,4xx,3xx,2xx,1xx
			// 9 -> CM, 8 -> DCCC, 7 -> DCC, 6 -> DC, 5 C, 4 CD, 3 CCC, 2 CC, 1 C
			roman += getDigit("C", "D", "M", num, 100)
			num %= 100
		} else if num >= 10 {
			roman += getDigit("X", "L", "C", num, 10)
			num %= 10
		} else if num >= 1 {
			roman += getDigit("I", "V", "X", num, 1)
			num %= 1
		} else {
			break
		}
	}
	return roman
}

func getDigit(oneValue, fiveValue, tenValue string, intValue int, size int) string {

	var roman string

	// we get 500
	remainder := intValue % size
	digit := (intValue - remainder) / size

	logrus.Infof("digit: %d", digit)

	// digit from 1 to 9

	if digit <= 3 {
		for i := 0; i < digit; i++ {
			roman += oneValue
		}
		return roman
	}

	if digit == 4 {
		roman += oneValue + fiveValue
		return roman
	}

	if digit == 5 {
		return fiveValue
	}

	if digit > 5 && digit < 9 {
		roman += fiveValue
		for i := 0; i < digit-5; i++ {
			roman += oneValue
		}
		return roman

	}

	if digit == 9 {
		return oneValue + tenValue
	}

	// C, M, D
	// 9 -> CM, 8 -> DCCC, 7 -> DCC, 6 -> DC, 5 C, 4 CD, 3 CCC, 2 CC, 1 C
	return roman
}
