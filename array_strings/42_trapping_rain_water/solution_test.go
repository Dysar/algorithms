package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		total := trap([]int{4, 2, 0, 3, 2, 5})
		require.Equal(t, 9, total)
	})

	t.Run("2", func(t *testing.T) {
		total := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		require.Equal(t, 6, total)
	})

}
