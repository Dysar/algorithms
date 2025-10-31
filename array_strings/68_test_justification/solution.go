package main

import (
	"fmt"
	"strings"
)

// You should pack your words in a greedy approach; that is, pack as many words as you can in each line.

// Extra spaces between words should be distributed as evenly as possible.
// If the number of spaces on a line does not divide evenly between words,
// the empty slots on the left will be assigned more spaces than the slots on the right.

/*

Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

Example2:
Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]

Example 3:
Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]

*/

func fullJustify(words []string, maxWidth int) []string {

	var (
		currentRowRemainingSpace = maxWidth
		rowWords                 = [][]string{
			{}, // init first element
		}
		currentRowIndex = 0
	)

	for _, word := range words {
		currentRowRemainingSpace -= len(word) + 1
		if currentRowRemainingSpace == -1 {
			// this means we can fit this word without the space. Do it.
			currentRowRemainingSpace = 0
		}

		fmt.Printf("remaining space after removed %s+1 (%d): %d (out of %d)\n", word, len(word)+1, currentRowRemainingSpace, maxWidth)
		if currentRowRemainingSpace < 0 {
			currentRowRemainingSpace = maxWidth - (len(word) + 1) // recalculate remaining space for next row
			currentRowIndex++                                     // move to next row
			rowWords = append(rowWords, []string{word})           // init next element with this word
		} else {
			rowWords[currentRowIndex] = append(rowWords[currentRowIndex], word) // add this word to current row
		}
	}

	fmt.Printf("row words: %s\n\n", rowWords)

	// deal with spaces
	finalRows := make([]string, len(rowWords))
	for i, row := range rowWords {

		if len(row) == 1 && i != len(rowWords)-1 { // 1 word in a row. Align left + spaces
			finalRows[i] += row[0]
			spacesToAdd := maxWidth - len(row[0]) //amount of spaces
			finalRows[i] += strings.Repeat(" ", spacesToAdd)
		} else if i == len(rowWords)-1 { // last row

			finalRows[i] = strings.Join(row, " ")
			fmt.Printf("last row: %s", row)
			spacesToAdd := maxWidth - len(finalRows[i]) //amount of spaces

			finalRows[i] += strings.Repeat(" ", spacesToAdd)
		} else { // normal case: multiple words, not last row.
			totalWordLen := 0

			for _, word := range row {

				totalWordLen += len(word)
			}
			totalSpaces := maxWidth - totalWordLen
			fmt.Printf("line (not last), (not 1 word): %s\n", row)
			fmt.Printf(" total words len %d, total spaces %d\n", totalWordLen, totalSpaces)

			spacesPerGap := totalSpaces / (len(row) - 1)
			remainder := totalSpaces % (len(row) - 1)
			fmt.Printf(" spaces per gap %d, remainder: %d\n", spacesPerGap, remainder)

			// spaces per gap * len(words)-1. E.g [4,4]
			spaces := make([]int, len(row)-1)
			for j := 0; j < len(row)-1; j++ {
				spaces[j] = spacesPerGap
			}
			for j := 1; j <= remainder; j++ { // 1,2,3
				spaces[j-1] += 1
			}

			for j := 0; j < len(row)-1; j++ {
				word := row[j]
				finalRows[i] += word + strings.Repeat(" ", spaces[j])
			}
			finalRows[i] += row[len(row)-1] // add last word

		}
	}

	return finalRows
}

func main() {
	result := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	fmt.Println()
	for _, res := range result {
		fmt.Printf("result row: %s\n", res)
		// fmt.Println(len(res))
	}
}
