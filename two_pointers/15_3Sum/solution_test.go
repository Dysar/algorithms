package _25_valid_palindrome

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := threeSum([]int{-1, 0, 1, 2, -1, -4})
		t.Log(output)
		require.Len(t, output, 2)
		for i, sum := range output {
			switch i {
			case 2:
				require.True(t, reflect.DeepEqual(sum, []int{-1, -1, 2}))
			case 1:
				require.True(t, reflect.DeepEqual(sum, []int{-1, 0, 1}))
			}
		}
	})
	t.Run("case 2", func(t *testing.T) {
		output := threeSum([]int{0, 1, 1})
		require.Len(t, output, 0)
	})
	t.Run("case 3", func(t *testing.T) {
		output := threeSum([]int{0, 0, 0})
		t.Log(output)
		require.True(t, reflect.DeepEqual(output[0], []int{0, 0, 0}))
	})

	t.Run("case 4", func(t *testing.T) {
		output := threeSum([]int{0, 0, 0, 0})
		t.Log(output)
		require.Len(t, output, 1)
		require.True(t, reflect.DeepEqual(output[0], []int{0, 0, 0}))
	})

}
