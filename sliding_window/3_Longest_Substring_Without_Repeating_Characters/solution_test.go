package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := lengthOfLongestSubstring("abcabcbb")
		require.Equal(t, 3, res)
	})
	t.Run("2", func(t *testing.T) {
		res := lengthOfLongestSubstring("bbbbb")
		require.Equal(t, 1, res)
	})
	t.Run("3", func(t *testing.T) {
		res := lengthOfLongestSubstring("pwwkew")
		require.Equal(t, 3, res)
	})
	t.Run("4", func(t *testing.T) {
		res := lengthOfLongestSubstring("aab")
		require.Equal(t, 2, res)
	})
}
