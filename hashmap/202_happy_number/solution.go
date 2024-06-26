package main

import "fmt"

func isHappy(n int) bool {
	sums := map[int]bool{}
	for {
		var digits []int
		divided := n
		if divided < 10 {
			digits = []int{n}
		}
		for divided >= 10 {
			divided = n / 10
			remainder := n % 10

			if divided < 10 {
				digits = append(digits, divided)
				if remainder != 0 {
					digits = append(digits, remainder)
				}
				break
			} else if remainder != 0 {
				digits = append(digits, remainder)
			}
			n = divided
			fmt.Println(n, digits)
		}

		sumSquares := 0
		for _, d := range digits {
			sumSquares += d * d
		}
		if _, ok := sums[sumSquares]; ok {
			return false
		} else {
			sums[sumSquares] = true
		}

		//fmt.Println(sumSquares, digits)

		if sumSquares == 1 {
			return true
		} else {
			n = sumSquares
		}
	}

	// 1 <= n <= 2pow31 - 1

}
