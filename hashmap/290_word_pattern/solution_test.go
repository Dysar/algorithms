package main

import "testing"

func Test_wordPattern(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		if ok := wordPattern("aaaa", "dog cat cat dog"); ok {
			t.Fatal("expected false")
		}
	})

	t.Run("case 2", func(t *testing.T) {
		if ok := wordPattern("aba", "dog cat dog"); !ok {
			t.Fatal("expected true")
		}
	})

	t.Run("case 3", func(t *testing.T) {
		if ok := wordPattern("abba", "dog dog dog dog"); ok {
			t.Fatal("expected false")
		}
	})

	t.Run("case 4", func(t *testing.T) {
		if ok := wordPattern("aaa", "aa aa aa aa"); ok {
			t.Fatal("expected false")
		}
	})
}
