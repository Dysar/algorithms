package main

import "testing"

func Test_canConstruct(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		if ok := canConstruct("a", "b"); ok {
			t.Fatal("expected false but got true")
		}
	})

	t.Run("case 2", func(t *testing.T) {
		if ok := canConstruct("aa", "ab"); ok {
			t.Fatal("expected false but got true")
		}
	})

	t.Run("case 3", func(t *testing.T) {
		if ok := canConstruct("aa", "aab"); !ok {
			t.Fatal("expected true but got false")
		}
	})

	t.Run("case 4", func(t *testing.T) {
		if ok := canConstruct("aab", "baa"); !ok {
			t.Fatal("expected true but got false")
		}
	})
}
