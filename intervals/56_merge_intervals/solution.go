package _6_merge_intervals

import (
	"sort"
)

// Given an array of intervals where intervals[i] = [starti, endi],
// merge all overlapping intervals, and return an array of the non-overlapping intervals
// that cover all the intervals in the input.
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { // O(nlogn)
		return intervals[i][0] < intervals[j][0] // Sort based on first element of inner slice
	})

	result := make([][]int, 0) // O(n) worst
	current := intervals[0]
	result = append(result, current)

	for _, interval := range intervals {
		//currBegin := current[0]
		currEnd := current[1]
		nextBegin := interval[0]
		nextEnd := interval[1]

		if currEnd >= nextBegin {
			// this will update it in the output
			current[1] = max(currEnd, nextEnd)
		} else {
			current = interval
			result = append(result, current)
		}
	}

	return result
}
