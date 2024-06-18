package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		if l := lengthOfLastWord("Hello World"); l != 5 {
			t.Error(l)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		if l := lengthOfLastWord("   fly me   to   the moon  "); l != 4 {
			t.Error(l)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		if l := lengthOfLastWord("luffy is still joyboy"); l != 6 {
			t.Error(l)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		if l := lengthOfLastWord("Today is a nice day"); l != 3 {
			t.Error(l)
		}
	})
}
