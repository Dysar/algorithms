package main

import (
	"fmt"
	"strings"
)

/* The string "PAYPALISHIRING" is written in a zigzag pattern
              "0123456789"

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
*/

func convert(s string, numRows int) string {
	// zig zag. top -> down. Diagonal up, top down

	// so if num rows 4

	// input: PAY PAL ISH IRI NG
	// output PIN ALS IGY AHR PI

	// indices: 0,1,2 3,4,5 6,7,8 9,10,11 12,13
	// output: 0,6,12 1,5,7

	// so if we go down and count; 1 2 3, turn, 1 2 3, turn, 1 2 3, turn then we have

	k := 0
	rows := make([][]string, numRows)
	isForwardDirection := true

	for _, char := range s {

		//1st row
		rows[k] = append(rows[k], string(char))

		fmt.Println(k, isForwardDirection)

		if numRows > 1 {

			if isForwardDirection {
				k++ // 0 ->1 -> 2-> 3-> 4; 1 row case -> k can be
				if k == numRows-1 {
					isForwardDirection = false
				}
			} else if !isForwardDirection {
				k-- // 3 -> 2; 2->1; 1->0
				if k == 0 {
					isForwardDirection = true
				}
			}

		}
	}

	// and now we should combine the rows together
	result := ""
	for _, row := range rows {
		result += strings.Join(row, "")
	}
	return result
}

func main() {

	/*
		P   A   H   N
		A P L S I I G
		Y   I   R
	*/
	fmt.Println(convert("PAYPALISHIRING", 3))
	// PAHNAPLSIIGYIR
}
