package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		s := "a good   example"
		result := reverseWords(s)
		require.Equal(t, "example good a", result)
	})
	t.Run("2", func(t *testing.T) {
		s := "the sky is blue"
		result := reverseWords(s)
		require.Equal(t, result, "blue is sky the")
	})
	t.Run("3", func(t *testing.T) {
		s := "  hello world  "
		result := reverseWords(s)
		require.Equal(t, result, "world hello")
	})

}
