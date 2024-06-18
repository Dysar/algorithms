package main

import "strings"

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	s = strings.TrimSpace(s)
	counter := 0

	// what cases can there be. Last word has spaces after it
	// last word is the first word
	for _, ch := range s {
		if string(ch) != " " {
			counter++
		} else {
			counter = 0
		}
	}
	return counter
}
