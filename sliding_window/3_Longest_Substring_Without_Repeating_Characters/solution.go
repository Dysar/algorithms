package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l, res := 0, 0
	seen := make(map[byte]bool)

	for r := 0; r < len(s); r++ {
		for seen[s[r]] { // shrink window until s[r] is unique
			delete(seen, s[l])
			l++
		}
		seen[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
