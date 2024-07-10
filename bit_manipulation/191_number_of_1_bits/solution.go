package main

import "fmt"

func hammingWeight(n int) int {
	res := 0
	for n != 0 {
		fmt.Println(n)
		if n&1 == 1 {
			res++
		}
		n >>= 1 // this is right shift, or diving by 2
	}

	// TC O(32) -> O(1)
	return res
}
