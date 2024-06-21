package main

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	//if len(strs) == 1 && strs[0] == "" {
	//	return [][]string{{""}}
	//}

	// the key is the count of characters. The value is a list of results.
	anagrams := make(map[[26]int][]string, len(strs))

	for i := 0; i < len(strs); i++ {

		count := [26]int{}
		for _, ch := range strs[i] {
			count[ch-'a']++
		}

		// if the counts are the same, append the result.
		if results, ok := anagrams[count]; ok {
			results = append(results, strs[i])
			anagrams[count] = results
		} else {
			//  If the results are not the same, create a new element
			anagrams[count] = []string{strs[i]}
		}
	}

	res := make([][]string, 0, len(anagrams))
	for _, anagram := range anagrams {
		res = append(res, anagram)
	}
	return res
}
