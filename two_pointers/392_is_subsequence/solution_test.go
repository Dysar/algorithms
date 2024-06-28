package main

import "testing"

func Test_isSubsequence(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := isSubsequence("abc", "ahbgdc")
		if output != true {
			t.Fatalf("expected: %v, got: %v", true, output)
		}
	})
	t.Run("case 1", func(t *testing.T) {
		output := isSubsequence("axc", "ahbgdc")
		if output != false {
			t.Fatalf("expected: %v, got: %v", false, output)
		}
	})
}
