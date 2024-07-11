package _25_valid_palindrome

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
		if output != 49 {
			t.Error("case 1 failed: ", output)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := maxArea([]int{1, 1})
		if output != 1 {
			t.Error("case 2 failed: ", output)
		}
	})

}
