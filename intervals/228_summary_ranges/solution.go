package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	answer := make([]string, 0, len(nums))
	// can solve without the store actually, just 2 indices, but ok
	var store []int

	if len(nums) == 0 {
		return answer
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1]+1 == nums[i] {
			store = append(store, nums[i])
			continue
		}
		answer = append(answer, makeString(store))
		store = nil
		store = append(store, nums[i])
	}
	return append(answer, makeString(store))
}

func makeString(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	if len(arr) == 1 {
		return fmt.Sprintf("%d", arr[0])
	}
	return fmt.Sprintf("%d->%d", arr[0], arr[len(arr)-1])
}
