package main

func canConstruct(ransomNote, magazine string) bool {

	if len(ransomNote) == 0 || len(magazine) == 0 {
		return false
	}

	charCount := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		charCount[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if count, ok := charCount[ransomNote[i]]; !ok || count <= 0 {
			return false
		} else {
			charCount[ransomNote[i]]--
		}
	}

	// TC: O(n+m); SC: O(n)
	return true
}
