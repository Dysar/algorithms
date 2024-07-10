package _5_search_insert_position

import "testing"

func Test_searchInsert(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		output := searchInsert([]int{1, 2, 3, 4, 5}, 4)
		if output != 3 {
			t.Errorf("output = %d, want %d", output, 3)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		output := searchInsert([]int{1, 3, 5, 6}, 2)
		if output != 1 {
			t.Errorf("output = %d, want %d", output, 1)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		output := searchInsert([]int{1, 3, 5, 6}, 7)
		if output != 4 {
			t.Errorf("output = %d, want %d", output, 4)
		}
	})
}
