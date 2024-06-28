package two_pointers

import "strings"

func isPalindrome(s string) bool {
	var newStr string
	for _, char := range s {
		if char >= 'a' && char <= 'z' || (char >= '0' && char <= '9') {
			newStr += string(char)
		}
		if char >= 'A' && char <= 'Z' {
			newStr += strings.ToLower(string(char))
		}
	}
	s = newStr

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
