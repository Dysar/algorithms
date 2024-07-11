package _25_valid_palindrome

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := twoSum([]int{2, 7, 11, 15}, 9)
		if output[0] != 1 || output[1] != 2 {
			t.Error("case 1 failed: ", output)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := twoSum([]int{2, 3, 4}, 6)
		if output[0] != 1 || output[1] != 3 {
			t.Error("case 2 failed: ", output)
		}
	})
	t.Run("case 3", func(t *testing.T) {
		output := twoSum([]int{-1, 0}, -1)
		if output[0] != 1 || output[1] != 2 {
			t.Error("case 3 failed: ", output)
		}
	})
	t.Run("case 4", func(t *testing.T) {
		output := twoSum([]int{-3, 24, 50, 79, 88, 150, 345}, 200)
		if output[0] != 3 || output[1] != 6 {
			t.Error("case 4 failed: ", output)
		}
	})

}
