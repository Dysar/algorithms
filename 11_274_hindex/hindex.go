package main

import (
	"sort"
)

// sample input citations = [3,0,6,1,5]
// means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
//
// The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
func hIndex(citations []int) int {

	// Sort in descending order: O(nlogn)
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	// c: [6,5,3,1,0]; [3,1,1]
	// i:  1,2,3,4,5;   1,2,3
	for i := 1; i <= len(citations); i++ { // O(n)
		if i == citations[i-1] {
			return i
		}
		if i > citations[i-1] {
			return i - 1
		}
	}

	return len(citations)
}
