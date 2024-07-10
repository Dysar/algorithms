package main

import "testing"

func Test_hammingWeight(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := hammingWeight(11)
		if output != 3 {
			t.Errorf("hammingWeight(11) returned %d", output)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := hammingWeight(128)
		if output != 1 {
			t.Errorf("hammingWeight(128) returned %d", output)
		}
	})

}
