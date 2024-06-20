package main

import "strings"

// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
//
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false
func wordPattern(pattern string, s string) bool {
	// pattern contains only lowercase english letters: 26 letters in total

	// a pattern char defines a numbers of chars in the result string.
	// SC O(26)
	mapping := make(map[string]string, 26)

	// TC O(n)
	ss := strings.Split(s, " ")

	if len(ss) != len(pattern) {
		return false
	}

	// TC O(n) where n is the pattern len (up to 300)
	for i := 0; i < len(pattern); i++ {
		word := ss[i]
		ch := string(pattern[i])

		key, ok := mapping[ch]
		// check if the current char already in the map
		if ok && key != word {
			// if there is already "a":"dog"
			return false
		} else if !ok {
			// we also need to check if there are duplicate values.
			for _, value := range mapping {
				if value == word {
					return false
				}
			}

			mapping[ch] = ss[i]
		}
	}

	return true
}
