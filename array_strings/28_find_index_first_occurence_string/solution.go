package main

// Input: haystack = "sadbutsad", needle = "sad"
func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			//if we found the potential match, we run a nested loop
			failed := false
			for j := 1; j < len(needle); j++ {

				// if the needle is longer than the haystack
				if len(haystack) == i+j {
					return -1
				}

				if haystack[i+j] != needle[j] {
					// if in this loop we find that it does not fully match the substring
					// we break to explore the next index
					failed = true
					break
				}
			}
			if !failed {
				return i
			}
		}
	}

	// SC O(1), TC O( n (len(haystack) * m (len(needle) )
	return -1
}
