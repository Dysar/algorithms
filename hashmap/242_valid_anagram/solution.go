package main

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sCharFreq := make(map[uint8]int)
	tCharFreq := make(map[uint8]int)

	for i := 0; i < len(s); i++ {

		//put s charFreqs
		if _, exists := sCharFreq[s[i]]; exists {
			sCharFreq[s[i]]++
		} else {
			sCharFreq[s[i]] = 1
		}

		// put t charFreqs
		if _, exists := tCharFreq[t[i]]; exists {
			tCharFreq[t[i]]++
		} else {
			tCharFreq[t[i]] = 1
		}
	}

	for key, sVal := range sCharFreq {
		tVal, ok := tCharFreq[key]
		if !ok {
			return false
		}
		if tVal != sVal {
			return false
		}
	}
	return true
}
