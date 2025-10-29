package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convert(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := convert("PAYPALISHIRING", 3)
		require.Equal(t, "PAHNAPLSIIGYIR", res)
	})

	t.Run("2", func(t *testing.T) {
		res := convert("PAYPALISHIRING", 4)
		require.Equal(t, "PINALSIGYAHRPI", res)
	})
	t.Run("3", func(t *testing.T) {
		res := convert("AB", 1)
		require.Equal(t, "AB", res)
	})
	t.Run("4", func(t *testing.T) {
		res := convert("ABCD", 2)
		require.Equal(t, "ACBD", res)
	})
}
