package main

import "testing"

func Test_addBinary(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := addBinary("11", "1")
		if output != "100" {
			t.Errorf("output was incorrect, got: %s, want: %s.", output, "100")
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := addBinary("1010", "1011")
		if output != "10101" {
			t.Errorf("output was incorrect, got: %s, want: %s.", output, "10101")
		}
	})
	t.Run("case 3", func(t *testing.T) {})
}
