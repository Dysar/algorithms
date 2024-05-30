package main

import "fmt"

var (
	// total count of beautiful subsets found
	totalBeautifulSubsets int
	// frequency array to keep track of elements in the sub
	count   = [1000]int{}
	theNums []int
	theK    int
	results [][]int
)

func main() {
	fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2))
	for i, c := range count {
		if c != 0 {
			fmt.Println(i, " ", c)
		}
	}
}

func beautifulSubsets(nums []int, k int) int {
	theNums = nums
	theK = k

	// start the DFS with the first element in the array
	findBeautifulSubsets(0)
	return totalBeautifulSubsets
}

func findBeautifulSubsets(index int) {
	// Base case: if we've considered all elements in 'theNums'
	if index >= len(theNums) {
		// we've formed a subset, increment the beautiful subsets count
		totalBeautifulSubsets++
		return
	}

	// Option 1: Exclude the current element and explore further subsets
	findBeautifulSubsets(index + 1)

	// Check if including theNums[index] in the subset still keeps it beautiful
	isBeautifulWithNext := theNums[index]+theK >= len(count) || count[theNums[index]+theK] == 0
	isBeautifulWithPrevious := theNums[index]-theK < 0 || count[theNums[index]-theK] == 0

	// If both checks pass, the subset remains beautiful with the inclusion of theNums[index]
	if isBeautifulWithNext && isBeautifulWithPrevious {
		// include the current element by incrementing its frequency
		count[theNums[index]]++
		findBeautifulSubsets(index + 1) // explore further subsets including this element
		count[theNums[index]]--         // backtrack and exclude the current element
	}
}
