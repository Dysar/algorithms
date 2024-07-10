package main

import "testing"

func Test_reverseBits(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := reverseBits(43261596)
		if output != 964176192 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 964176192)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := reverseBits(4294967293)
		if output != 3221225471 {
			t.Errorf("Output was incorrect, got: %d, want: %d.", output, 3221225471)
		}
	})
}
