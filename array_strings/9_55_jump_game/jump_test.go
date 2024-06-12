package main

import "testing"

func Test_jump(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		case1 := []int{2, 3, 1, 1, 4}
		if !canJump2(case1) {
			t.Error("case 1 failed")
		}
	})

	t.Run("case 2", func(t *testing.T) {
		case2 := []int{3, 2, 1, 0, 4}
		if canJump2(case2) {
			t.Error("case 2 failed")
		}
	})
}
