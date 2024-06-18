package main

import "testing"

func Test_strStr(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		if index := strStr("sadbutsad", "sad"); index != 0 {
			t.Error(index)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		if index := strStr("leetcode", "leeto"); index != -1 {
			t.Error(index)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		if index := strStr("mississippi", "issip"); index != 4 {
			t.Error(index)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		if index := strStr("aaa", "aaaa"); index != -1 {
			t.Error(index)
		}
	})

}
