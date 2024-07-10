package main

import "testing"

func Test_singleNumber(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := singleNumber([]int{2, 2, 1})
		if output != 1 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 964176192)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := singleNumber([]int{4, 1, 2, 1, 2})
		if output != 4 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 3221225471)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		output := singleNumber([]int{1})
		if output != 1 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 964176192)
		}
	})
}
