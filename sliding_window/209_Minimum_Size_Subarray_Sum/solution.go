package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	l, total := 0, 0
	minCnt := len(nums) + 1
	for r := 0; r < len(nums); r++ {
		total += nums[r]
		for total >= target {
			minCnt = min(minCnt, r-l+1)
			total -= nums[l]
			l++
		}
	}
	if minCnt == len(nums)+1 {
		return 0
	}
	return minCnt
}
func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
