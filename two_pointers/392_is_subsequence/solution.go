package main

func isSubsequence(s string, t string) bool {
	sIdx, tIdx := 0, 0

	// example for s =abc, t = ahbgdc
	// sIdx [0,1,1,2,2,3
	// tIdx [0,1,2,3,4,5
	for sIdx < len(s) && tIdx < len(t) {
		//a = a
		if s[sIdx] == t[tIdx] {
			sIdx++
			tIdx++
		} else {
			tIdx++
		}
	}
	return sIdx == len(s)
}
