package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSubArrayLen(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
		require.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		res := minSubArrayLen(4, []int{1, 4, 4})
		require.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		res := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
		require.Equal(t, 0, res)
	})
	t.Run("4", func(t *testing.T) {
		res := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
		require.Equal(t, 3, res)
	})
}
