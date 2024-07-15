package main

import "testing"

func Test_numIslands(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		one := byte('1')
		zero := byte('0')
		grid := [][]byte{
			{one, one, one, one, zero},
			{one, one, zero, one, zero},
			{one, one, zero, zero, zero},
			{zero, zero, zero, zero, zero},
		}
		output := numIslands(grid)
		if output != 1 {
			t.Errorf("case 1 failed. got %d; want %d", output, 1)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		one := byte('1')
		zero := byte('0')
		grid := [][]byte{
			{one, one, zero, zero, zero},
			{one, one, zero, zero, zero},
			{zero, zero, one, zero, zero},
			{zero, zero, zero, one, one},
		}
		output := numIslands(grid)
		if output != 3 {
			t.Errorf("case 2 failed. got %d; want %d", output, 3)
		}
	})
}
