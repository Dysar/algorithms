package main

import (
	"fmt"
)

/*

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

*/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	} else if t == s {
		return s
	} else if len(t) == 0 {
		return ""
	}

	checkMap := map[rune]int{}

	// we need to make a map for each element in string t
	for _, ch := range t {
		checkMap[ch] = 0
	}

	have, need := 0, len(checkMap)
	resLen := len(s) // the best len
	indices := []int{}
	l := 0
	bestWindow := ""

	// so we need to traverse the strings s and use a sliding map here. We can use 2 pointers for this.
	for r := 0; r < len(s); r++ {
		charRightPtr := s[r]
		if val, ok := checkMap[rune(charRightPtr)]; ok {
			indices = append(indices, r)
			if val == 0 {
				have++ // currently we increment have for any char. but we should do it for the missing char only
			}
			checkMap[rune(charRightPtr)]++ // incr value
		}
		if have == need {

			resLen = min(resLen, r-l+1)

			have--
			hasSomeZeroVal := false
			for !hasSomeZeroVal {
				// we take the left index and decrement the counter in the map
				l = indices[0]
				val := s[l]
				checkMap[rune(val)]--
				// we must make sure the checkMap has some 0 value
				for _, val := range checkMap {
					if val == 0 {
						hasSomeZeroVal = true
					}
				}

				indices = indices[1:] // remove left-most index from our window
			}

			// get the window
			localBestWindow := ""
			for i := l; i < r+1; i++ {
				localBestWindow += string(s[i])
			}
			bestWindow = localBestWindow

			fmt.Printf("index: %d, got all 3 chars, removed from the map 1 char, best len:%d\n", r, resLen)
		}

		fmt.Println(l, r, checkMap, have, indices, bestWindow)
	}

	return bestWindow

}

func main() {
	fmt.Print(minWindow("bbaa", "aba"))
}
