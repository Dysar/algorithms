package main

import "fmt"

/*

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

*/

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	// make max left arr
	// make max right arr
	// make min(L,R) and calc water on the fyl

	maxL := []int{}
	var maxLeftVal int
	for i := 0; i < len(height); i++ { // 0,1,2
		if i == 0 {
			maxL = append(maxL, 0)
		} else {
			leftVal := height[i-1] // 1,2
			if leftVal > maxLeftVal {
				maxLeftVal = leftVal
			}
			maxL = append(maxL, maxLeftVal)
		}
	}

	maxR := make([]int, len(height))
	var maxRightVal int
	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			maxR[i] = 0 // no value at the right of last.
		} else {
			rightVal := height[i+1]
			if rightVal > maxRightVal {
				maxRightVal = rightVal
			}
			maxR[i] = maxRightVal
		}
	}

	fmt.Println(maxL)
	fmt.Println(maxR)

	totalWater := 0
	for i := 0; i < len(height); i++ {
		if val := min(maxL[i], maxR[i]) - height[i]; val > 0 {
			totalWater += val
		}
	}

	return totalWater
}

func main() {
	trap([]int{1, 2, 3})
}
