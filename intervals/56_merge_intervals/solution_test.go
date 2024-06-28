package _6_merge_intervals

import "testing"

func Test_merge(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		intervals := merge([][]int{
			{1, 3}, {2, 6}, {8, 10}, {15, 18},
		})
		if len(intervals) != 3 {
			t.Log(intervals)
			t.Fatalf("%d", len(intervals))
		}
	})
	t.Run("case 2", func(t *testing.T) {
		intervals := merge([][]int{
			{1, 4}, {4, 5},
		})
		if len(intervals) != 1 {
			t.Fail()
		}
		if intervals[0][0] != 1 {
			t.Fail()
		}
		if intervals[0][1] != 5 {
			t.Fail()
		}
		t.Log(intervals)
	})

	t.Run("case 3", func(t *testing.T) {
		intervals := merge([][]int{
			{4, 5}, {1, 4},
		})
		if len(intervals) != 1 {
			t.Fail()
		}
		if intervals[0][0] != 1 {
			t.Fail()
		}
		if intervals[0][1] != 5 {
			t.Fail()
		}
		t.Log(intervals)
	})

	t.Run("case 4", func(t *testing.T) {
		intervals := merge([][]int{
			{1, 4}, {5, 6},
		})
		if len(intervals) != 2 {
			t.Fail()
		}
		if intervals[0][0] != 1 {
			t.Fail()
		}
		if intervals[1][1] != 6 {
			t.Fail()
		}
		t.Log(intervals)
	})

	t.Run("case 5 multiple overlapping", func(t *testing.T) {
		intervals := merge([][]int{
			{1, 4}, {0, 2}, {3, 5},
		})
		if len(intervals) != 1 {
			t.Fail()
		}
		t.Log(intervals)
	})
}
