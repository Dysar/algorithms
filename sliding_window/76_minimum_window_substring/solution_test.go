package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minWindow(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := minWindow("ADOBECODEBANC", "ABC")
		require.Equal(t, "BANC", res)
	})

	t.Run("2", func(t *testing.T) {
		res := minWindow("a", "a")
		require.Equal(t, "a", res)
	})

	t.Run("3", func(t *testing.T) {
		res := minWindow("a", "aa")
		require.Equal(t, "", res)
	})

	t.Run("4", func(t *testing.T) {
		res := minWindow("abc", "ab")
		require.Equal(t, "ab", res)
	})

	t.Run("5", func(t *testing.T) {
		res := minWindow("bbaa", "aba")
		require.Equal(t, "baa", res)
	})
}
