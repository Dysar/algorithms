package main

import "testing"

func Test_jump(t *testing.T) {
	t.Run("classic", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		minJumps := jump(nums)
		if minJumps != 2 {
			t.Errorf("calculation failed, expected 2, got %d", minJumps)
		}
	})

	t.Run("3 elements", func(t *testing.T) {
		nums := []int{1, 2, 3}
		minJumps := jump(nums)
		if minJumps != 2 {
			t.Errorf("calculation failed, expected 2, got %d", minJumps)
		}
	})
}
