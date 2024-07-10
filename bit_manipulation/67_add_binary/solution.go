package main

import (
	"strconv"
	"strings"
)

// Given two binary strings a and b, return their sum as a binary string.
func addBinary(a string, b string) string {

	// Input: a = "11", b = "1"
	//Output: "100"
	//  11
	// 	1
	//  ==
	// 100
	sb := strings.Builder{}
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 {

		sum := carry
		if i >= 0 {
			sum += int(a[i] - uint8('0')) // 1-0 = 1, 0-0 = 0
		}
		if j >= 0 {
			sum += int(b[j] - uint8('0'))
		}
		sb.Write([]byte(strconv.Itoa(sum % 2))) // if 1+1=2 we carry 1 and write 0. if 1+0 we write 0
		carry = sum / 2

		i--
		j-- // if 2 -> 1; if 1 -> 0
	}

	if carry != 0 {
		sb.Write([]byte(strconv.Itoa(carry)))
	}
	return reverse(sb.String())
}

func reverse(s string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}
