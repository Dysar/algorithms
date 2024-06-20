package main

import "testing"

func Test_isAnagram(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := isAnagram("anagram", "nagaram")
		if output != true {
			t.Fatal()
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := isAnagram("rat", "car")
		if output != false {
			t.Fatal()
		}
	})

	t.Run("case 3", func(t *testing.T) {
		output := isAnagram("a", "ab")
		if output != false {
			t.Fatal()
		}
	})
}
