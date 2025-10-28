package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intToRoman(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		num := 3749
		want := "MMMDCCXLIX"
		got := intToRoman(num)
		require.Equal(t, want, got)
	})

	t.Run("case 2", func(t *testing.T) {
		num := 58
		want := "LVIII"
		got := intToRoman(num)
		require.Equal(t, want, got)
	})
}

func Test_getDigit(t *testing.T) {

	digit := getDigit("M", "D", "C", 3999, 1000)
	require.Equal(t, "MMM", digit)

	digit = getDigit("C", "D", "M", 399, 100)
	require.Equal(t, "CCC", digit)

	digit = getDigit("X", "L", "C", 39, 10)
	require.Equal(t, "XXX", digit)

	digit = getDigit("I", "V", "X", 9, 1)
	require.Equal(t, "IX", digit)
}

func Test_testModulo(t *testing.T) {
	require.Equal(t, 999, 3999%1000)
	require.Equal(t, 99, 399%100)
	require.Equal(t, 9, 39%10)
	require.Equal(t, 0, 3%1)
}
