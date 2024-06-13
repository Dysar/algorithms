package main

import "testing"

func Test_plusOne(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		digits := []int{1, 2, 3}
		if res := plusOne(digits); res[len(res)-1] != 4 {
			t.Errorf("plusOne returned %v, expected last to be 4", res)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		digits := []int{9}
		if res := plusOne(digits); res[1] != 0 || res[0] != 1 {
			t.Errorf("plusOne returned %v, expected res to be [1,0]", res)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		digits := []int{9, 9}
		if res := plusOne(digits); res[len(res)-1] != 0 {
			t.Log(res)
			t.Errorf("plusOne returned %v, expected last to be 0", res)
		}
	})
}
