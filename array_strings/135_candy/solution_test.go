package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_candy(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := candy([]int{1, 0, 2})
		require.Equal(t, 5, res)
	})
	t.Run("2", func(t *testing.T) {
		res := candy([]int{1, 2, 2})
		require.Equal(t, 4, res)
	})
}
