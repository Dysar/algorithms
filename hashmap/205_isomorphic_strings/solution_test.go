package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		res := isIsomorphic("egg", "add")
		if res != true {
			t.Fatalf("got %v want %v", res, true)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		res := isIsomorphic("foo", "bar")
		if res != false {
			t.Fatalf("got %v want %v", res, false)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		res := isIsomorphic("paper", "title")
		if res != true {
			t.Fatalf("got %v want %v", res, true)
		}
	})
}
